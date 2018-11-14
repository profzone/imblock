package account_service

import (
	"github.com/profzone/imblock/core"
	"github.com/profzone/imblock/core/message_bus"
	"github.com/profzone/imblock/core/castle"
	"github.com/sirupsen/logrus"
	"fmt"
	"github.com/profzone/imblock/global"
	"github.com/profzone/imblock/account_service/base"
	"github.com/johnnyeven/libtools/courier"
	"github.com/profzone/imblock/account_service/base/api/accounts"
)

var _ interface {
	castle.Service
} = (*AccountBaseServiceBootstrap)(nil)

type AccountBaseServiceBootstrap struct {
	service AccountService
}

func NewAccountBaseServiceBootstrap() castle.Service {
	return &AccountBaseServiceBootstrap{
		service: base.NewBaseBootstrap(),
	}
}

func (s *AccountBaseServiceBootstrap) Messages() []message_bus.MessageHandler {
	return []message_bus.MessageHandler{
		{
			Topic: message_bus.TOPIC_CREATE_ACCOUNT,
			Runner: func(message message_bus.Message) (reply message_bus.Message, err error) {
				data, ok := message.Data.(map[string]string)
				if !ok {
					err = fmt.Errorf("[AccountBaseServiceBootstrap] Handle message error: topic=%s, err=%s", message_bus.TOPIC_CREATE_ACCOUNT, "request parameter is not map")
					return
				}
				account := s.service.CreateAccount(data["alias"])
				if account == nil {
					err = fmt.Errorf("[AccountBaseServiceBootstrap] Handle message error: topic=%s, err=%s", message_bus.TOPIC_CREATE_ACCOUNT, "alias has been used")
					return
				}
				reply.Topic = message.Topic
				reply.Data = map[string]string{
					"alias":   account.GetAlias(),
					"address": string(account.GetAddress()),
					"pubKey":  fmt.Sprintf("0x%x", account.GetPubKey()),
				}
				return
			},
		},
		{
			Topic: message_bus.TOPIC_GET_ACCOUNTS,
			Runner: func(message message_bus.Message) (reply message_bus.Message, err error) {
				accounts := s.service.GetAccounts()
				reply.Topic = message.Topic

				replies := make([]map[string]string, 0)
				for _, account := range accounts {
					replies = append(replies, map[string]string{
						"alias":   account.GetAlias(),
						"address": string(account.GetAddress()),
						"pubKey":  fmt.Sprintf("0x%x", account.GetPubKey()),
					})
				}
				reply.Data = replies
				return
			},
		},
	}
}

func (s *AccountBaseServiceBootstrap) Protocols() []core.ProtocolHandler {
	return nil
}

func (s *AccountBaseServiceBootstrap) Routes() []*courier.Router {
	return []*courier.Router{
		accounts.Router,
	}
}

func (s *AccountBaseServiceBootstrap) Start() error {
	err := s.service.Load(global.GetWalletFilePath())
	if err != nil {
		logrus.Errorf("[AccountBaseServiceBootstrap] start failed, err: %v", err)
		return err
	}
	logrus.Info("[AccountBaseServiceBootstrap] started.")
	return nil
}

func (s *AccountBaseServiceBootstrap) Stop() error {
	err := s.service.Save(global.GetWalletFilePath())
	if err != nil {
		logrus.Errorf("[AccountBaseServiceBootstrap] stop failed, err: %v", err)
		return err
	}
	logrus.Info("[AccountBaseServiceBootstrap] stopped.")
	return nil
}
