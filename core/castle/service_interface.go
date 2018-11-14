package castle

import (
	"github.com/profzone/imblock/core/message_bus"
	"github.com/profzone/imblock/core"
	"github.com/johnnyeven/libtools/courier"
)

type Service interface {
	Messages() []message_bus.MessageHandler
	Protocols() []core.ProtocolHandler
	Routes() []*courier.Router
	Start() error
	Stop() error
}

type ServiceConstructor func() Service
