package api_service

import (
	"github.com/profzone/imblock/api_service/serializable"
)

type ApiService interface {

	// Peers
	GetPeers() ([]serializable.Peer, error)

	// Accounts
	CreateAccount(alias string) (serializable.Account, error)
	GetAccounts() ([]serializable.Account, error)
}
