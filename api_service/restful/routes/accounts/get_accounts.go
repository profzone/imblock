package accounts

import (
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"context"
	"github.com/profzone/imblock/api_service/restful/services"
)

func init() {
	Router.Register(courier.NewRouter(GetAccounts{}))
}

// 获取账户列表
type GetAccounts struct {
	httpx.MethodGet
}

func (req GetAccounts) Path() string {
	return ""
}

func (req GetAccounts) Output(ctx context.Context) (result interface{}, err error) {
	return services.GetApi().GetAccounts()
}

