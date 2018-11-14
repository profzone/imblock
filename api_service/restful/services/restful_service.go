package services

import (
	"github.com/profzone/imblock/core"
	"github.com/profzone/imblock/api_service/serializable"
	"fmt"
	"github.com/johnnyeven/libtools/timelib"
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
