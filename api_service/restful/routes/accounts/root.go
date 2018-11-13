package accounts

import (
	"github.com/johnnyeven/libtools/courier"
)

var Router = courier.NewRouter(Accounts{})

// 账户管理接口
type Accounts struct {
	courier.EmptyOperator
}

func (req Accounts) Path() string {
	return "/accounts"
}
