package peers

import (
	"context"
	"fmt"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/libtools/timelib"
	"github.com/profzone/imblock/api_service/model"
	"github.com/profzone/imblock/core"
)

func init() {
	Router.Register(courier.NewRouter(GetPeers{}))
}

// 获取节点列表
type GetPeers struct {
	httpx.MethodGet
}

func (req GetPeers) Path() string {
	return ""
}

func (req GetPeers) Output(ctx context.Context) (result interface{}, err error) {
	replies := make([]model.Peer, 0)
	core.GetPeerManager().Iterator(func(peer *core.Peer) error {

		replies = append(replies, model.Peer{
			Guid: fmt.Sprintf("0x%x", peer.Guid),
			Node: model.Node{
				ID:             peer.Node.ID.HexString(),
				Addr:           peer.Node.Addr.String(),
				LastActiveTime: timelib.MySQLDatetime(peer.Node.LastActiveTime).String(),
			},
			Heartbeat: model.Heartbeat{
				Delay:  peer.Heartbeat.Delay,
				Health: peer.Heartbeat.Health,
			},
		})
		return nil
	}, false)

	return replies, nil
}
