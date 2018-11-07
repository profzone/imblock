package memory

import (
	"testing"
	"github.com/derekparker/trie"
	"github.com/profzone/imblock/core/message_bus"
	"reflect"
	"github.com/profzone/data_structures/list"
	"fmt"
)

var m = MessageDriverMemory{
	tree: trie.New(),
}

var testCase = map[string][]message_bus.MessageRunner{
	"foo": {
		func(message message_bus.Message) (message_bus.Message, error) {
			return message_bus.Message{
				Data: fmt.Sprintf("%s, this is foo handler1 result", message.Data),
			}, nil
		},
		func(message message_bus.Message) (message_bus.Message, error) {
			return message_bus.Message{
				Data: fmt.Sprintf("%s, this is foo handler2 result", message.Data),
			}, nil
		},
	},
}

func TestSubscribe(t *testing.T) {

	for topic, handlers := range testCase {
		for _, handler := range handlers {
			m.Subscribe(topic, handler)
		}
	}

	for topic := range testCase {
		if node, ok := m.tree.Find(topic); ok {
			meta := node.Meta()
			if handlers, ok := meta.(*list.SyncedList); ok {
				for e := range handlers.Iter() {
					handler := e.Value.(message_bus.MessageRunner)
					result, _ := handler(message_bus.Message{
						Data: "TestSubscribe",
					})
					t.Log(result.Data)
				}
			} else {
				t.Errorf("Subscribe faild, expected *list.SyncedList, got %v", reflect.TypeOf(meta).String())
			}
		} else {
			t.Errorf("Subscribe faild, expected %v, got nil", topic)
		}
	}
}

func TestPublish(t *testing.T) {
	result, err := m.Publish(message_bus.Message{
		Topic: "foo",
		Data:  "TestPublish",
	})

	if err != nil {
		t.Fatal(err)
		return
	}

	if len(result) != 2 {
		t.Fatalf("len(result) expected 2, got %d", len(result))
	}

	for _, message := range result {
		t.Log(message.Data)
	}
}

func TestAsyncPublish(t *testing.T) {
	msg := message_bus.Message{
		Topic: "foo",
		Data:  "TestAsyncPublish",
	}

	result, err := m.AsyncPublish(msg)
	if err != nil {
		t.Fatal(err)
		return
	}

	for reply := range result {
		t.Log(reply.Data)
	}

}

func TestUnsubscribe(t *testing.T) {
	m.Unsubscribe("foo", testCase["foo"][0])

	if node, ok := m.tree.Find("foo"); ok {
		meta := node.Meta()
		if handlers, ok := meta.(*list.SyncedList); ok {
			length := handlers.Len()
			if length != 1 {
				t.Errorf("handlers.Len() expected 1, got %d", length)
			}
			for e := range handlers.Iter() {
				handler := e.Value.(message_bus.MessageRunner)
				result, _ := handler(message_bus.Message{
					Data: "TestUnsubscribe",
				})
				t.Log(result.Data)
			}
		} else {
			t.Errorf("Unsubscribe faild, expected *list.SyncedList, got %v", reflect.TypeOf(meta).String())
		}
	}
}
