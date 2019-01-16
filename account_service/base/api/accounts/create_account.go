package accounts

import (
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"context"
	"github.com/profzone/imblock/account_service/base"
	"fmt"
	"github.com/profzone/imblock/account_service/base/model"
	"github.com/profzone/imblock/core/constant"
)

func init() {
	Router.Register(courier.NewRouter(CreateAccount{}))
}

type CreateAccountBody struct {
	Alias string `json:"alias"`
}

// 创建账户
type CreateAccount struct {
	httpx.MethodPost

	Body CreateAccountBody `in:"body" name:"body"`
}

func (req CreateAccount) Path() string {
	return ""
}

func (req CreateAccount) Output(ctx context.Context) (result interface{}, err error) {
	account := base.GetBase().CreateAccount(req.Body.Alias)

	if account == nil {
		err = fmt.Errorf("[AccountBaseServiceBootstrap] Handle message error: topic=%s, err=%s", constant.TOPIC_ACCOUNT_CREATE, "alias has been used")
		return
	}

	result = model.BaseAccount{
		Alias:   account.GetAlias(),
		Address: string(account.GetAddress()),
		PubKey:  fmt.Sprintf("0x%x", account.GetPubKey()),
	}

	return
}
