package peers

import (
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"context"
	"github.com/profzone/imblock/api_service/restful/services"
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
	return services.GetApi().GetPeers()
}
