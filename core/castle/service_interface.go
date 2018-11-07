package castle

import (
	"github.com/profzone/imblock/core/message_bus"
	"github.com/profzone/imblock/core"
)

type Service interface {
	Messages() []message_bus.MessageHandler
	Protocols() []core.ProtocolHandler
	Start() error
	Stop() error
}

type ServiceConstructor func() Service
