package game

import (
	"context"
	"errors"
	"fmt"
	"hash/crc32"

	pbGlobal "github.com/east-eden/server/proto/global"
	"github.com/east-eden/server/services/game/player"
	"github.com/east-eden/server/transport"
	"github.com/east-eden/server/utils"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ErrUnregistedMsgName = errors.New("unregisted message name")
)

func (m *MsgRegister) handleAccountTest(ctx context.Context, sock transport.Socket, p *transport.Message) error {
	return nil
}

func (m *MsgRegister) handleWaitResponseMessage(ctx context.Context, sock transport.Socket, p *transport.Message) error {
	msg, ok := p.Body.(*pbGlobal.C2S_WaitResponseMessage)
	if !ok {
		return errors.New("handleWaitResponseMessage failed: cannot assert value to message")
	}

	handler, err := m.r.GetHandler(msg.GetInnerMsgCrc())
	if !utils.ErrCheck(err, "handleWaitResponseMessage GetHandler by MsgCrc failed", msg.GetInnerMsgCrc()) {
		return ErrUnregistedMsgName
	}

	var innerMsg transport.Message
	innerMsg.Name = handler.Name
	innerMsg.Body, err = sock.PbMarshaler().Unmarshal(msg.GetInnerMsgData(), handler.RType)
	if !utils.ErrCheck(err, "handleWaitResponseMessage protobuf Unmarshal failed") {
		return err
	}

	// direct handle inner message
	err = handler.Fn(ctx, sock, &innerMsg)
	if !utils.ErrCheck(err, "handle inner message failed", handler.Name) {
		return err
	}

	err = m.am.AddAccountTask(
		ctx,
		m.am.GetAccountIdBySock(sock),
		func(c context.Context, p ...interface{}) error {
			acct := p[0].(*player.Account)
			m := p[1].(*pbGlobal.C2S_WaitResponseMessage)
			reply := &pbGlobal.S2C_WaitResponseMessage{
				MsgId:   m.MsgId,
				ErrCode: 0,
			}

			acct.SendProtoMessage(reply)
			return nil
		},
		msg,
	)

	return err
}

func (m *MsgRegister) handleAccountPing(ctx context.Context, sock transport.Socket, p *transport.Message) error {
	msg, ok := p.Body.(*pbGlobal.C2S_Ping)
	if !ok {
		return errors.New("handleAccountLogon failed: cannot assert value to message")
	}

	reply := &pbGlobal.S2C_Pong{
		Pong: msg.Ping + 1,
	}

	var send transport.Message
	send.Name = string(reply.ProtoReflect().Descriptor().Name())
	send.Body = reply

	return sock.Send(&send)
}

func (m *MsgRegister) handleAccountLogon(ctx context.Context, sock transport.Socket, p *transport.Message) error {
	msg, ok := p.Body.(*pbGlobal.C2S_AccountLogon)
	if !ok {
		return errors.New("handleAccountLogon failed: cannot assert value to message")
	}

	// todo userid暂时为crc32
	userId := crc32.ChecksumIEEE([]byte(msg.UserId))
	err := m.am.Logon(ctx, int64(userId), sock)
	if err != nil {
		return fmt.Errorf("handleAccountLogon failed: %w", err)
	}

	return err
}

func (m *MsgRegister) handleHeartBeat(ctx context.Context, sock transport.Socket, p *transport.Message) error {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
		m.timeHistogram.WithLabelValues("handleHeartBeat").Observe(v)
	}))

	_, ok := p.Body.(*pbGlobal.C2S_HeartBeat)
	if !ok {
		return errors.New("handleHeartBeat failed: cannot assert value to message")
	}

	err := m.am.AddAccountTask(
		ctx,
		m.am.GetAccountIdBySock(sock),
		func(c context.Context, p ...interface{}) error {
			acct := p[0].(*player.Account)
			defer timer.ObserveDuration()
			acct.HeartBeat()
			return nil
		},
	)

	return err
}

// client disconnect
func (m *MsgRegister) handleAccountDisconnect(ctx context.Context, sock transport.Socket, p *transport.Message) error {
	return player.ErrAccountDisconnect
}
