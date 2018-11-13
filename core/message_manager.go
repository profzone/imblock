package core

import (
	"github.com/profzone/imblock/core/message_bus"
	"github.com/profzone/imblock/core/message_bus/memory"
)

var messageManager *message_bus.MessageBus

func init() {
	messageManager = memory.CreateMemoryMessageBus()
}

func GetMessageManager() *message_bus.MessageBus {
	return messageManager
}