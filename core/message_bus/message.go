package message_bus

const (
	TOPIC_CREATE_ACCOUNT = "account/create"
	TOPIC_GET_ACCOUNTS = "account/list"
)

type Message struct {
	TransID   int64
	ServiceID int64
	Topic     string
	Data      map[string]interface{}
}
