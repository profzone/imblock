package account_service

import "github.com/profzone/imblock/core/model"

type AccountService interface {
	CreateAccount(alias string) model.Account
	GetAccount(address []byte) model.Account
	GetAccounts() []model.Account

	AccountStorage
}

type AccountStorage interface {
	Save(filePath string) error
	Load(filePath string) error
}