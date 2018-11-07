package memory

import (
	"github.com/profzone/imblock/core/message_bus"
	"github.com/derekparker/trie"
	"github.com/profzone/data_structures/list"
	"reflect"
	"errors"
	"fmt"
)

var _ interface {
	message_bus.MessageDriver
} = (*MessageDriverMemory)(nil)

type MessageDriverMemory struct {
	tree *trie.Trie
}

func (m *MessageDriverMemory) Subscribe(topic string, handler message_bus.MessageRunner) {
	var handlers *list.SyncedList
	if node, ok := m.tree.Find(topic); ok {
		handlers, _ = node.Meta().(*list.SyncedList)
	} else {
		handlers = list.NewSyncedList()
	}
	handlers.PushBack(handler)
	m.tree.Add(topic, handlers)
}

func (m *MessageDriverMemory) Unsubscribe(topic string, handler message_bus.MessageRunner) {
	var handlers *list.SyncedList
	if node, ok := m.tree.Find(topic); ok {
		handlers, _ = node.Meta().(*list.SyncedList)
	} else {
		return
	}

	if handlers.Len() == 0 {
		return
	}

	for e := range handlers.Iter() {
		if reflect.ValueOf(e.Value).Pointer() == reflect.ValueOf(handler).Pointer() {
			handlers.Remove(e)
		}
	}

	if handlers.Len() == 0 {
		m.tree.Remove(topic)
	} else {
		m.tree.Add(topic, handlers)
	}
}

func (m *MessageDriverMemory) Publish(message message_bus.Message) (result []message_bus.Message, err error) {
	if message.Topic == "" {
		return nil, errors.New("topic must NOT be empty")
	}

	topics := m.tree.PrefixSearch(message.Topic)
	if len(topics) == 0 {
		return nil, errors.New(fmt.Sprintf("can NOT find any topics by topic(prefix) %s", message.Topic))
	}

	for _, topic := range topics {

		if node, ok := m.tree.Find(topic); ok {

			handlers, _ := node.Meta().(*list.SyncedList)
			for e := range handlers.Iter() {
				handler := e.Value.(message_bus.MessageRunner)
				handlerResult, err := handler(message)
				if err != nil {
					return nil, err
				}

				result = append(result, handlerResult)
			}
		} else {
			return nil, errors.New(fmt.Sprintf("can NOT find any handler by topic %s", topic))
		}
	}

	return
}

func (m *MessageDriverMemory) AsyncPublish(message message_bus.Message) (<-chan message_bus.Message, error) {
	if message.Topic == "" {
		return nil, errors.New("topic must NOT be empty")
	}

	topics := m.tree.PrefixSearch(message.Topic)
	if len(topics) == 0 {
		return nil, errors.New(fmt.Sprintf("can NOT find any topics by topic(prefix) %s", message.Topic))
	}

	replyChannel := make(chan message_bus.Message)

	go func() {
		for _, topic := range topics {

			if node, ok := m.tree.Find(topic); ok {

				handlers, _ := node.Meta().(*list.SyncedList)
				for e := range handlers.Iter() {
					handler := e.Value.(message_bus.MessageRunner)
					handlerResult, _ := handler(message)

					replyChannel <- handlerResult
				}
			}
		}
		close(replyChannel)
	}()

	return replyChannel, nil
}

func CreateMemoryMessageBus() *message_bus.MessageBus {
	driver := &MessageDriverMemory{
		tree: trie.New(),
	}
	return message_bus.CreateMessageBus(driver)
}
