package services

import (
	"github.com/profzone/imblock/core"
	"github.com/profzone/imblock/api_service/serializable"
	"fmt"
	"github.com/johnnyeven/libtools/timelib"
	"github.com/profzone/imblock/core/message_bus"
	"github.com/profzone/imblock/account_service"
)

var api *HttpService

func init() {
	api = &HttpService{}
}

func GetApi() *HttpService {
	return api
}

type HttpService struct{}

func (*HttpService) GetPeers() (result []serializable.Peer, err error) {
	result = make([]serializable.Peer, 0)
	core.GetPeerManager().Iterator(func(peer *core.Peer) error {

		result = append(result, serializable.Peer{
			Guid: fmt.Sprintf("0x%x", peer.Guid),
			Node: serializable.Node{
				ID:             peer.Node.ID.HexString(),
				Addr:           peer.Node.Addr.String(),
				LastActiveTime: timelib.MySQLDatetime(peer.Node.LastActiveTime).String(),
			},
			Heartbeat: serializable.Heartbeat{
				Delay:  peer.Heartbeat.Delay,
				Health: peer.Heartbeat.Health,
			},
		})
		return nil
	}, false)

	return
}

func (*HttpService) CreateAccount(alias string) (result serializable.Account, err error) {
	request := message_bus.Message{
		Topic: message_bus.TOPIC_CREATE_ACCOUNT,
		Data: map[string]interface{}{
			"alias": alias,
		},
	}
	replies, err := core.GetMessageManager().Publish(request)
	if err != nil {
		return
	}

	msg := replies[0]
	account, _ := msg.Data["account"].(account_service.Account)
	result.Alias = account.GetAlias()
	result.Address = string(account.GetAddress())
	result.PubKey = fmt.Sprintf("0x%x", account.GetPubKey())

	return
}

func (*HttpService) GetAccounts() (result []serializable.Account, err error) {
	request := message_bus.Message{
		Topic: message_bus.TOPIC_GET_ACCOUNTS,
	}
	replies, err := core.GetMessageManager().Publish(request)
	if err != nil {
		return
	}

	msg := replies[0]
	accounts, _ := msg.Data["accounts"].([]account_service.Account)

	for _, account := range accounts {
		result = append(result, serializable.Account{
			Alias:   account.GetAlias(),
			Address: string(account.GetAddress()),
			PubKey:  fmt.Sprintf("0x%x", account.GetPubKey()),
		})
	}

	return
}
