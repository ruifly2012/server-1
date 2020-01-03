package game

import (
	"context"
	"log"
	"sync"

	logger "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/yokaiio/yokai_server/game/player"
	"github.com/yokaiio/yokai_server/internal/define"
	"github.com/yokaiio/yokai_server/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type PlayerManager struct {
	g                     *Game
	listPlayer            map[int64]*player.Player
	listAccountPlayer     map[int64](map[int64]*player.Player)
	listLitePlayer        map[int64]*player.LitePlayer
	listAccountLitePlayer map[int64](map[int64]*player.LitePlayer)

	chExpire     chan int64
	chLiteExpire chan int64
	wg           utils.WaitGroupWrapper
	ctx          context.Context
	cancel       context.CancelFunc
	coll         *mongo.Collection
	sync.RWMutex
}

func NewPlayerManager(g *Game, ctx *cli.Context) *PlayerManager {
	m := &PlayerManager{
		g:                     g,
		listPlayer:            make(map[int64]*player.Player, 0),
		listAccountPlayer:     make(map[int64](map[int64]*player.Player), 0),
		listLitePlayer:        make(map[int64]*player.LitePlayer, 0),
		listAccountLitePlayer: make(map[int64](map[int64]*player.LitePlayer), 0),
		chExpire:              make(chan int64, define.Player_ExpireChanNum),
		chLiteExpire:          make(chan int64, define.Player_ExpireChanNum),
	}

	m.ctx, m.cancel = context.WithCancel(ctx)

	// migrate
	m.migrate()

	return m
}

func (m *PlayerManager) TableName() string {
	return "player"
}

func (m *PlayerManager) migrate() {
	m.coll = m.g.ds.Database().Collection(m.TableName())

	// create index
	_, err := m.coll.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bsonx.Doc{{"account_id", bsonx.Int32(1)}},
		},
	)

	if err != nil {
		logger.Warn("player manager create index failed:", err)
	}

	//player.Migrate(ds)
}

func (m *PlayerManager) addPlayer(p *player.Player) {
	m.Lock()
	defer m.Unlock()

	// map id to player
	m.listPlayer[p.GetID()] = p

	// map account_id to player list
	listPlayer, ok := m.listAccountPlayer[p.GetAccountID()]
	if !ok {
		listPlayer = make(map[int64]*player.Player, 0)
		m.listAccountPlayer[p.GetAccountID()] = listPlayer
	}

	if _, ok := listPlayer[p.GetID()]; ok {
		delete(listPlayer, p.GetID())
	}

	listPlayer[p.GetID()] = p
}

func (m *PlayerManager) removePlayer(id int64) {
	m.Lock()
	if p, ok := m.listPlayer[id]; ok {
		if mapPlayer, ok := m.listAccountPlayer[p.GetAccountID()]; ok {
			delete(mapPlayer, id)
		}
	}
	delete(m.listPlayer, id)
	m.Unlock()
}

func (m *PlayerManager) beginTimeExpire(p *player.Player) {
	// memcache time expired
	go func() {
		for {
			select {
			case <-m.ctx.Done():
				return

			case <-p.GetExpire().C:
				// if account online, then reset memcache time expire
				if acct := m.g.am.GetAccountByID(p.GetAccountID()); acct != nil {
					p.ResetExpire()
					continue
				}

				m.chExpire <- p.GetID()
			}
		}
	}()
}

func (m *PlayerManager) addLitePlayer(p *player.LitePlayer) {
	m.Lock()
	defer m.Unlock()

	// map id to player
	m.listLitePlayer[p.GetID()] = p

	// map account_id to player list
	listPlayer, ok := m.listAccountLitePlayer[p.GetAccountID()]
	if !ok {
		listPlayer = make(map[int64]*player.LitePlayer, 0)
		m.listAccountLitePlayer[p.GetAccountID()] = listPlayer
	}

	if _, ok := listPlayer[p.GetID()]; ok {
		delete(listPlayer, p.GetID())
	}

	listPlayer[p.GetID()] = p
}

func (m *PlayerManager) removeLitePlayer(id int64) {
	m.Lock()
	if p, ok := m.listLitePlayer[id]; ok {
		if mapPlayer, ok := m.listAccountLitePlayer[p.GetAccountID()]; ok {
			delete(mapPlayer, id)
		}
	}
	delete(m.listLitePlayer, id)
	m.Unlock()
}

func (m *PlayerManager) beginLiteTimeExpire(p *player.LitePlayer) {
	// memcache time expired
	go func() {
		for {
			select {
			case <-m.ctx.Done():
				return

			case <-p.GetExpire().C:
				m.chLiteExpire <- p.GetID()
			}
		}
	}()
}

// slow load player lite info, first hit in memory, then find in database
func (m *PlayerManager) slowLoadLitePlayers(accountID int64) map[int64]*player.LitePlayer {
	res := m.coll.FindOne(m.ctx, bson.D{{"account_id", accountID}})
	if res.Err() == nil {
		p := player.NewLitePlayer(-1, "")
		res.Decode(p)
		m.addLitePlayer(p)
		m.beginLiteTimeExpire(p)
	}

	return map[int64]*player.LitePlayer{}
}

// slow load player, first hit in memory, then find in database
func (m *PlayerManager) slowLoadPlayers(accountID int64) map[int64]*player.Player {
	res := m.coll.FindOne(m.ctx, bson.D{{"account_id", accountID}})
	if res.Err() == nil {
		p := player.NewPlayer(-1, "", m.g.ds)
		res.Decode(p)
		m.addPlayer(p)
		m.beginTimeExpire(p)
	}

	return m.listAccountPlayer[accountID]
}

// slow load player lite info, first hit in memory, then find in database
func (m *PlayerManager) slowLoadPlayerByID(playerID int64) *player.LitePlayer {
	resp, err := m.g.rpcHandler.CallGetLitePlayer(playerID)
	if err != nil {
		logger.WithFields(logger.Fields{
			"player_id": playerID,
			"err":       err.Error(),
		}).Warn("slow load player error")
		return nil
	}

	return &player.LitePlayer{
		ID:        resp.Info.Id,
		AccountID: resp.Info.AccountId,
		Name:      resp.Info.Name,
		Exp:       resp.Info.Exp,
		Level:     resp.Info.Level,
	}
}

func (m *PlayerManager) Main() error {
	exitCh := make(chan error)
	var once sync.Once
	exitFunc := func(err error) {
		once.Do(func() {
			if err != nil {
				log.Fatal("PlayerManager Main() error:", err)
			}
			exitCh <- err
		})
	}

	m.wg.Wrap(func() {
		exitFunc(m.Run())
	})

	return <-exitCh
}

func (m *PlayerManager) Run() error {
	for {
		select {
		case <-m.ctx.Done():
			logger.Print("player manager context done!")
			return nil

		// memcache time expired
		case id := <-m.chExpire:
			m.removePlayer(id)

			// lite player memcache time expired
		case id := <-m.chLiteExpire:
			m.removeLitePlayer(id)
		}

	}

	return nil
}

func (m *PlayerManager) Exit() {
	logger.Info("PlayerManager context done...")
	m.cancel()
	m.wg.Wait()
}

func (m *PlayerManager) GetLitePlayers(accountID int64) map[int64]*player.LitePlayer {
	m.RLock()
	mapPlayer := m.listAccountPlayer[accountID]
	m.RUnlock()

	if len(mapPlayer) > 0 {
		mapLitePlayer := make(map[int64]*player.LitePlayer)
		for _, v := range mapPlayer {
			v.ResetExpire()
			mapLitePlayer[v.GetID()] = v.LitePlayer
		}

		return mapLitePlayer
	}

	return m.slowLoadLitePlayers(accountID)
}

func (m *PlayerManager) GetPlayerByID(id int64) *player.Player {
	m.RLock()
	p, ok := m.listPlayer[id]
	m.RUnlock()

	if ok {
		p.ResetExpire()
	} else {
		res := m.coll.FindOne(m.ctx, bson.D{{"_id", id}})
		if res.Err() == nil {
			p = player.NewPlayer(id, "", m.g.ds)
			res.Decode(p)
			m.addPlayer(p)
			m.beginTimeExpire(p)
		}
	}

	return p
}

func (m *PlayerManager) GetPlayersByAccountID(accountID int64) map[int64]*player.Player {
	m.RLock()
	mapPlayer := m.listAccountPlayer[accountID]
	m.RUnlock()

	for _, v := range mapPlayer {
		v.ResetExpire()
		return mapPlayer
	}

	return m.slowLoadPlayers(accountID)
}

func (m *PlayerManager) GetOnePlayerByAccountID(accountID int64) *player.Player {
	m.RLock()
	mapPlayers := m.listAccountPlayer[accountID]
	m.RUnlock()

	for _, v := range mapPlayers {
		v.ResetExpire()
		return v
	}

	res := m.coll.FindOne(m.ctx, bson.D{{"account_id", accountID}})
	if res.Err() == nil {
		p := player.NewPlayer(-1, "", m.g.ds)
		res.Decode(p)
		m.addPlayer(p)
		m.beginTimeExpire(p)
		return p
	}

	return nil
}

func (m *PlayerManager) CreatePlayer(accountID int64, name string) (*player.Player, error) {
	id, err := utils.NextID(define.Plugin_Player)
	if err != nil {
		return nil, err
	}

	p := player.NewPlayer(id, name, m.g.ds)
	p.SetAccountID(accountID)
	p.Save()

	m.addPlayer(p)
	m.beginTimeExpire(p)

	return p, nil
}

func (m *PlayerManager) ExpirePlayer(playerID int64) {
	m.RLock()
	_, ok := m.listPlayer[playerID]
	m.RUnlock()

	if ok {
		m.chExpire <- playerID
	}
}

func (m *PlayerManager) ExpireLitePlayer(playerID int64) {
	m.RLock()
	_, ok := m.listLitePlayer[playerID]
	m.RUnlock()

	if ok {
		m.chLiteExpire <- playerID
	}
}
