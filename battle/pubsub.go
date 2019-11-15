package battle

import (
	"context"

	"github.com/micro/go-micro"
	logger "github.com/sirupsen/logrus"
	pbClient "github.com/yokaiio/yokai_server/proto/client"
	pbPubSub "github.com/yokaiio/yokai_server/proto/pubsub"
)

type PubSub struct {
	pubBattleResult micro.Publisher
	b               *Battle
}

func NewPubSub(b *Battle) *PubSub {
	ps := &PubSub{
		b: b,
	}

	// create publisher
	ps.pubBattleResult = micro.NewPublisher("battle.BattleResult", b.mi.srv.Client())

	// register subscriber
	micro.RegisterSubscriber("game.StartBattle", b.mi.srv.Server(), &subStartBattle{b: b})

	return ps
}

/////////////////////////////////////
// publish handle
/////////////////////////////////////
func (ps *PubSub) PubBattleResult(ctx context.Context, win bool) error {
	info := &pbClient.ClientInfo{Id: 1, Name: "pub_client"}
	return ps.pubBattleResult.Publish(ps.b.ctx, &pbPubSub.PubBattleResult{Info: info, Win: win})
}

/////////////////////////////////////
// subscribe handle
/////////////////////////////////////
type subStartBattle struct {
	b *Battle
}

func (s *subStartBattle) Process(ctx context.Context, event *pbPubSub.PubStartBattle) error {
	logger.WithFields(logger.Fields{
		"event": event,
	}).Info("recv game.StartBattle")
	return nil
}