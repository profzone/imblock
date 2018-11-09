package api_service

import (
	"github.com/profzone/imblock/api_service/serializable"
)

type ApiService interface {
	GetPeers() ([]serializable.Peer, error)
}
