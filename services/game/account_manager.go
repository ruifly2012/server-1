package game

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"bitbucket.org/funplus/server/define"
	"bitbucket.org/funplus/server/services/game/player"
	"bitbucket.org/funplus/server/services/game/prom"
	"bitbucket.org/funplus/server/store"
	"bitbucket.org/funplus/server/transport"
	"bitbucket.org/funplus/server/utils"
	"bitbucket.org/funplus/server/utils/cache"
	"github.com/golang/groupcache/lru"
	"github.com/golang/protobuf/proto"
	log "github.com/rs/zerolog/log"

	// "github.com/sasha-s/go-deadlock"
	"github.com/urfave/cli/v2"
)

var (
	maxPlayerInfoLruCache = 10000            // max number of lite player, expire non used PlayerInfo
	AccountCacheExpire    = 10 * time.Minute // 账号cache缓存10分钟

	ErrAccountHasNoPlayer = errors.New("account has no player")
	ErrAccountNotFound    = errors.New("account not found")
)

type AccountManagerFace interface {
}

type AccountManager struct {
	cacheAccounts *cache.Cache
	mapSocks      map[transport.Socket]int64 // socket->accountId

	g  *Game
	wg utils.WaitGroupWrapper

	accountConnectMax int

	playerPool      sync.Pool
	accountPool     sync.Pool
	playerInfoPool  sync.Pool
	playerInfoCache *lru.Cache

	sync.RWMutex
}

func NewAccountManager(ctx *cli.Context, g *Game) *AccountManager {
	am := &AccountManager{
		g:                 g,
		cacheAccounts:     cache.New(AccountCacheExpire, AccountCacheExpire),
		mapSocks:          make(map[transport.Socket]int64),
		accountConnectMax: ctx.Int("account_connect_max"),
		playerInfoCache:   lru.New(maxPlayerInfoLruCache),
	}

	// 账号缓存删除时处理
	am.cacheAccounts.OnEvicted(func(k, v interface{}) {
		acct := v.(*player.Account)

		am.Lock()
		delete(am.mapSocks, acct.GetSock())
		am.Unlock()

		acct.Close()
		am.playerPool.Put(acct.GetPlayer())
		am.accountPool.Put(v)
		log.Info().Interface("key", k).Msg("account cache evicted")
	})

	am.playerPool.New = player.NewPlayer
	am.accountPool.New = player.NewAccount
	am.playerInfoPool.New = player.NewPlayerInfo
	am.playerInfoCache.OnEvicted = am.OnPlayerInfoEvicted

	// add store info
	store.GetStore().AddStoreInfo(define.StoreType_Account, "account", "_id")
	store.GetStore().AddStoreInfo(define.StoreType_Player, "player", "_id")
	store.GetStore().AddStoreInfo(define.StoreType_Item, "player_item", "_id")
	store.GetStore().AddStoreInfo(define.StoreType_Hero, "player_hero", "_id")
	store.GetStore().AddStoreInfo(define.StoreType_Token, "player_token", "_id")
	store.GetStore().AddStoreInfo(define.StoreType_Fragment, "player_fragment", "_id")

	// migrate users table
	if err := store.GetStore().MigrateDbTable("account", "user_id"); err != nil {
		log.Fatal().Err(err).Msg("migrate collection account failed")
	}

	// migrate player table
	if err := store.GetStore().MigrateDbTable("player", "account_id"); err != nil {
		log.Fatal().Err(err).Msg("migrate collection player failed")
	}

	// migrate item table
	if err := store.GetStore().MigrateDbTable("player_item", "owner_id"); err != nil {
		log.Fatal().Err(err).Msg("migrate collection player_item failed")
	}

	// migrate hero table
	if err := store.GetStore().MigrateDbTable("player_hero", "owner_id"); err != nil {
		log.Fatal().Err(err).Msg("migrate collection player_hero failed")
	}

	// migrate hero table
	if err := store.GetStore().MigrateDbTable("player_token", "owner_id"); err != nil {
		log.Fatal().Err(err).Msg("migrate collection player_token failed")
	}

	// migrate fragment table
	if err := store.GetStore().MigrateDbTable("player_fragment", "owner_id"); err != nil {
		log.Fatal().Err(err).Msg("migrate collection player_fragment failed")
	}

	log.Info().Msg("AccountManager init ok ...")
	return am
}

func (am *AccountManager) OnPlayerInfoEvicted(key lru.Key, value interface{}) {
	am.playerInfoPool.Put(value)
}

func (am *AccountManager) Main(ctx context.Context) error {
	exitCh := make(chan error)
	var once sync.Once
	exitFunc := func(err error) {
		once.Do(func() {
			if err != nil {
				log.Fatal().Err(err).Msg("AccountManager Main() failed")
			}
			exitCh <- err
		})
	}

	am.wg.Wrap(func() {
		defer utils.CaptureException()
		exitFunc(am.Run(ctx))
	})

	return <-exitCh
}

func (am *AccountManager) Exit() {
	am.wg.Wait()
	log.Info().Msg("account manager exit...")
}

func (am *AccountManager) loadPlayer(acct *player.Account) error {
	ids := acct.GetPlayerIDs()
	if len(ids) < 1 {
		return ErrAccountHasNoPlayer
	}

	p := am.playerPool.Get().(*player.Player)
	p.Init()
	p.SetAccount(acct)
	err := store.GetStore().FindOne(context.Background(), define.StoreType_Player, ids[0], p)
	if !utils.ErrCheck(err, "load player object failed", ids[0]) {
		am.playerPool.Put(p)
		return err
	}

	// 加载玩家其他数据
	err = p.AfterLoad()
	if !utils.ErrCheck(err, "player.AfterLoad failed", ids[0]) {
		am.playerPool.Put(p)
		return err
	}

	acct.SetPlayer(p)

	return nil
}

func (am *AccountManager) handleLoadPlayer(ctx context.Context, acct *player.Account, msg *transport.Message) error {
	err := am.loadPlayer(acct)

	// 还没有角色
	if errors.Is(err, ErrAccountHasNoPlayer) {
		acct.SendLogon()
		return nil
	}

	if err == nil {
		// send logon success
		acct.SendLogon()

		// sync to client
		acct.GetPlayer().SendInitInfo()

		// 时间跨度检查
		acct.GetPlayer().CheckTimeChange()
	}

	return err
}

// 踢掉account对象
func (am *AccountManager) KickAccount(ctx context.Context, acctId int64, gameId int32) error {
	if acctId == -1 {
		return nil
	}

	// 踢掉本服account
	if int16(gameId) != am.g.ID {
		finish := make(chan int, 1)

		// stop account run
		err := am.AccountLazyHandle(acctId, &player.AccountLazyHandler{
			F: func(ctx context.Context, acct *player.Account, msg *transport.Message) error {
				close(finish)
				return player.ErrAccountKicked
			},
			M: nil,
		})

		// account 不在线
		if errors.Is(err, ErrAccountNotFound) {
			return nil
		}

		// 超时
		select {
		case <-ctx.Done():
			return errors.New("kick account timeout")
		case <-finish:
			break
		}

		// account.Run 结束后会自动删除account对象
		return nil

	} else {
		// game节点不存在的话不用发送rpc
		nodeId := fmt.Sprintf("game-%d", gameId)
		srvs, err := am.g.mi.srv.Server().Options().Registry.GetService("game")
		if err != nil {
			return nil
		}

		hit := false
		for _, srv := range srvs {
			for _, node := range srv.Nodes {
				if node.Id == nodeId {
					hit = true
					break
				}
			}
		}

		if !hit {
			return nil
		}

		// 发送rpc踢掉其他服account
		rs, err := am.g.rpcHandler.CallKickAccountOffline(acctId, gameId)
		if !utils.ErrCheck(err, "kick account offline failed", acctId, gameId, rs) {
			return err
		}

		// rpc调用成功
		if rs.GetAccountId() == acctId && rs.GetErrorCode() == 0 {
			return nil
		}

		return errors.New("kick account invalid error")
	}
}

func (am *AccountManager) addAccount(ctx context.Context, userId int64, accountId int64, accountName string, sock transport.Socket) error {
	if accountId == -1 {
		return errors.New("AccountManager.addAccount failed: account id invalid!")
	}

	am.RLock()
	socksNum := len(am.mapSocks)
	am.RUnlock()
	if socksNum >= am.accountConnectMax {
		return errors.New("AccountManager.addAccount failed: Reach game server's max account connect num")
	}

	acct := am.accountPool.Get().(*player.Account)
	acct.Init()
	acct.SetRpcCaller(am.g.rpcHandler)

	err := store.GetStore().FindOne(context.Background(), define.StoreType_Account, accountId, acct)
	if err != nil && !errors.Is(err, store.ErrNoResult) {
		return fmt.Errorf("AccountManager.addAccount failed: %w", err)
	}

	// 如果account的上次登陆game节点不是此节点，则发rpc提掉上一个登陆节点的account
	if acct.GameId != -1 && acct.GameId != am.g.ID {
		err := am.KickAccount(ctx, acct.ID, int32(acct.GameId))
		if !utils.ErrCheck(err, "kick account failed", acct.ID, acct.GameId, am.g.ID) {
			return err
		}
	}

	if errors.Is(err, store.ErrNoResult) {
		// 账号首次登陆
		acct.ID = accountId
		acct.UserId = userId
		acct.GameId = am.g.ID
		acct.Name = accountName

		// save object
		if err := store.GetStore().UpdateOne(context.Background(), define.StoreType_Account, acct.ID, acct); err != nil {
			log.Warn().
				Int64("account_id", accountId).
				Int64("user_id", userId).
				Err(err).
				Msg("save account failed")
		}
	} else {
		// 更新account节点id
		fields := map[string]interface{}{
			"game_id": acct.GameId,
		}

		err := store.GetStore().UpdateFields(context.Background(), define.StoreType_Account, acct.ID, fields)
		if !utils.ErrCheck(err, "account save game_id failed", acct.ID, acct.GameId) {
			return err
		}
	}

	// add account to manager
	am.Lock()
	am.cacheAccounts.Set(acct.GetID(), acct, AccountCacheExpire)
	am.mapSocks[sock] = acct.GetID()
	am.Unlock()

	acct.SetSock(sock)

	// load player
	err = am.AccountLazyHandle(acct.ID, &player.AccountLazyHandler{
		F: am.handleLoadPlayer,
		M: nil,
	})

	log.Info().
		Int64("user_id", acct.UserId).
		Int64("account_id", acct.ID).
		Str("name", acct.GetName()).
		Str("socket_remote", acct.GetSock().Remote()).
		Msg("add account success")

	// account run
	am.wg.Wrap(func() {
		defer utils.CaptureException()

		err := acct.Run(ctx)
		utils.ErrPrint(err, "account run failed", acct.GetID())

		// 记录下线时间
		fields := map[string]interface{}{
			"last_logoff_time": acct.LastLogoffTime,
		}
		err = store.GetStore().UpdateFields(context.Background(), define.StoreType_Account, acct.ID, fields)
		utils.ErrPrint(err, "account save last_logoff_time failed", acct.ID, acct.LastLogoffTime)

		// 删除缓存
		am.cacheAccounts.Delete(acct.GetID())
	})

	// prometheus ops
	prom.OpsOnlineAccountGauge.Set(float64(am.cacheAccounts.ItemCount()))
	prom.OpsLogonAccountCounter.Inc()

	return err
}

func (am *AccountManager) AccountLogon(ctx context.Context, userID int64, accountID int64, accountName string, sock transport.Socket) error {
	k, ok := am.cacheAccounts.Get(accountID)

	// if reconnect with same socket, then do nothing
	if ok && k.(*player.Account).GetSock() == sock {
		return nil
	}

	// if reconnect with another socket, replace socket in account
	if ok {
		acct := k.(*player.Account)
		if acct.GetSock() != nil {
			acct.GetSock().Close()
		}

		am.Lock()
		am.mapSocks[sock] = acct.GetID()
		am.Unlock()

		acct.SetSock(sock)
	}

	// add a new account with socket
	return am.addAccount(ctx, userID, accountID, accountName, sock)
}

func (am *AccountManager) GetAccountIdBySock(sock transport.Socket) int64 {
	am.RLock()
	defer am.RUnlock()

	return am.mapSocks[sock]
}

func (am *AccountManager) GetAccountById(acctId int64) *player.Account {
	acct, ok := am.cacheAccounts.Get(acctId)
	if ok {
		return acct.(*player.Account)
	}

	return nil
}

// add handler to account's execute channel, will be dealed by account's run goroutine
func (am *AccountManager) AccountLazyHandle(acctId int64, handler *player.AccountLazyHandler) error {
	acct := am.GetAccountById(acctId)

	if acct == nil {
		return ErrAccountNotFound
	}

	acct.LazyHandler <- handler
	return nil
}

func (am *AccountManager) CreatePlayer(acct *player.Account, name string) (*player.Player, error) {
	// only can create one player
	if pl, _ := am.GetPlayerByAccount(acct); pl != nil {
		return nil, player.ErrCreateMoreThanOnePlayer
	}

	id, err := utils.NextID(define.SnowFlake_Player)
	if err != nil {
		return nil, err
	}

	p := am.playerPool.Get().(*player.Player)
	p.Init()
	p.AccountID = acct.ID
	p.SetAccount(acct)
	p.SetID(id)
	p.SetName(name)

	// save handle
	errHandle := func(f func() error) {
		if err != nil {
			return
		}

		err = f()
	}
	errHandle(func() error {
		return store.GetStore().UpdateOne(context.Background(), define.StoreType_Player, p.ID, p)
	})

	errHandle(func() error {
		return store.GetStore().UpdateOne(context.Background(), define.StoreType_Hero, p.ID, p.HeroManager())
	})

	errHandle(func() error {
		return store.GetStore().UpdateOne(context.Background(), define.StoreType_Item, p.ID, p.ItemManager())
	})

	errHandle(func() error {
		return store.GetStore().UpdateOne(context.Background(), define.StoreType_Token, p.ID, p.TokenManager())
	})

	errHandle(func() error {
		return store.GetStore().UpdateOne(context.Background(), define.StoreType_Fragment, p.ID, p.FragmentManager())
	})

	// 保存失败处理
	if !utils.ErrCheck(err, "save player failed when CreatePlayer", id, name) {
		am.playerPool.Put(p)
		return nil, err
	}

	acct.SetPlayer(p)
	acct.Name = name
	acct.Level = p.GetLevel()
	acct.AddPlayerID(p.GetID())
	if err := store.GetStore().UpdateOne(context.Background(), define.StoreType_Account, acct.ID, acct); err != nil {
		log.Warn().
			Int64("account_id", acct.ID).
			Int64("user_id", acct.UserId).
			Err(err).
			Msg("save account failed")
	}

	// 同步玩家初始信息
	p.SendInitInfo()

	// todo 第一次上线处理

	return p, err
}

func (am *AccountManager) GetPlayerByAccount(acct *player.Account) (*player.Player, error) {
	if acct == nil {
		return nil, errors.New("invalid account")
	}

	if p := acct.GetPlayer(); p != nil {
		return p, nil
	}

	return nil, errors.New("invalid player")
}

func (am *AccountManager) GetPlayerInfo(playerId int64) (player.PlayerInfo, error) {
	am.RLock()
	defer am.RUnlock()

	if lp, ok := am.playerInfoCache.Get(playerId); ok {
		return *(lp.(*player.PlayerInfo)), nil
	}

	lp := am.playerInfoPool.Get().(*player.PlayerInfo)
	lp.Init()
	err := store.GetStore().FindOne(context.Background(), define.StoreType_Player, playerId, lp)
	if err == nil {
		am.playerInfoCache.Add(lp.ID, lp)
		return *lp, nil
	}

	am.playerInfoPool.Put(lp)
	return player.PlayerInfo{}, err
}

func (am *AccountManager) BroadCast(msg proto.Message) {
	items := am.cacheAccounts.Items()
	for _, v := range items {
		acct := v.Object.(*player.Account)

		acct.LazyHandler <- &player.AccountLazyHandler{
			F: func(ctx context.Context, a *player.Account, p *transport.Message) error {
				a.SendProtoMessage(msg)
				return nil
			},
		}
	}
}

func (am *AccountManager) Run(ctx context.Context) error {
	<-ctx.Done()
	log.Info().Msg("world session context done...")
	return nil
}
