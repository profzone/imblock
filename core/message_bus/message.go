package message_bus

type Message struct {
	TransID   uint64
	ServiceID uint64
	Topic     string
	Data      interface{}
}
