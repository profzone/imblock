package accounts

import (
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"context"
	"github.com/profzone/imblock/api_service/restful/services"
)

func init() {
	Router.Register(courier.NewRouter(CreateAccount{}))
}

type CreaAccountBody struct {
	Alias string `json:"alias"`
}

// 创建账户
type CreateAccount struct {
	httpx.MethodPost

	Body CreaAccountBody `in:"body" name:"body"`
}

func (req CreateAccount) Path() string {
	return ""
}

func (req CreateAccount) Output(ctx context.Context) (result interface{}, err error) {
	return services.GetApi().CreateAccount(req.Body.Alias)
}
