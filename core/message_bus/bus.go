package message_bus

var _ interface {
	MessageDriver
} = (*MessageBus)(nil)

type MessageBus struct {
	driver MessageDriver
}

func (bus *MessageBus) Subscribe(topic string, handler MessageRunner) {
	bus.driver.Subscribe(topic, handler)
}

func (bus *MessageBus) Unsubscribe(topic string, handler MessageRunner) {
	bus.driver.Unsubscribe(topic, handler)
}

func (bus *MessageBus) Publish(message Message) ([]Message, error) {
	return bus.driver.Publish(message)
}

func (bus *MessageBus) AsyncPublish(message Message) (<-chan Message, error) {
	return bus.driver.AsyncPublish(message)
}

func CreateMessageBus(driver MessageDriver) *MessageBus {
	return &MessageBus{
		driver: driver,
	}
}
