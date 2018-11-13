package account_service

import (
	"github.com/profzone/imblock/core"
	"github.com/profzone/imblock/core/message_bus"
	"github.com/profzone/imblock/core/castle"
	"github.com/sirupsen/logrus"
	"fmt"
	"github.com/profzone/imblock/global"
)

var _ interface {
	castle.Service
} = (*AccountB58ServiceBootstrap)(nil)

type AccountB58ServiceBootstrap struct {
	service AccountService
}

func NewAccountB58ServiceBootstrap() castle.Service {
	return &AccountB58ServiceBootstrap{
		service: NewB58Bootstrap(),
	}
}

func (s *AccountB58ServiceBootstrap) Messages() []message_bus.MessageHandler {
	return []message_bus.MessageHandler{
		{
			Topic: message_bus.TOPIC_CREATE_ACCOUNT,
			Runner: func(message message_bus.Message) (reply message_bus.Message, err error) {
				alias, ok := message.Data["alias"].(string)
				if !ok {
					err = fmt.Errorf("[AccountB58ServiceBootstrap] Handle message error: topic=%s, err=%s", message_bus.TOPIC_CREATE_ACCOUNT, "request parameter alias is not string")
					return
				}
				account := s.service.CreateAccount(alias)
				if account == nil {
					err = fmt.Errorf("[AccountB58ServiceBootstrap] Handle message error: topic=%s, err=%s", message_bus.TOPIC_CREATE_ACCOUNT, "alias has been used")
					return
				}
				reply.Topic = message.Topic
				reply.Data = map[string]interface{}{
					"account": account,
				}
				return
			},
		},
		{
			Topic: message_bus.TOPIC_GET_ACCOUNTS,
			Runner: func(message message_bus.Message) (reply message_bus.Message, err error) {
				accounts := s.service.GetAccounts()
				reply.Topic = message.Topic
				reply.Data = map[string]interface{}{
					"accounts": accounts,
				}
				return
			},
		},
	}
}

func (s *AccountB58ServiceBootstrap) Protocols() []core.ProtocolHandler {
	return nil
}

func (s *AccountB58ServiceBootstrap) Start() error {
	err := s.service.Load(global.GetWalletFilePath())
	if err != nil {
		logrus.Errorf("[AccountB58ServiceBootstrap] start failed, err: %v", err)
		return err
	}
	logrus.Info("[AccountB58ServiceBootstrap] started.")
	return nil
}

func (s *AccountB58ServiceBootstrap) Stop() error {
	err := s.service.Save(global.GetWalletFilePath())
	if err != nil {
		logrus.Errorf("[AccountB58ServiceBootstrap] stop failed, err: %v", err)
		return err
	}
	logrus.Info("[AccountB58ServiceBootstrap] stopped.")
	return nil
}
