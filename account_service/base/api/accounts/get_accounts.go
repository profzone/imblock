package accounts

import (
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"context"
	"github.com/profzone/imblock/account_service/base"
	"github.com/profzone/imblock/account_service/base/model"
	"fmt"
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
	accounts := base.GetBase().GetAccounts()
	replies := make([]model.BaseAccount, 0)

	for _, account := range accounts {
		replies = append(replies, model.BaseAccount{
			Alias:   account.GetAlias(),
			Address: string(account.GetAddress()),
			PubKey:  fmt.Sprintf("0x%x", account.GetPubKey()),
		})
	}
	return replies, nil
}

