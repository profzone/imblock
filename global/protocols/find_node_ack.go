package protocols

import (
	"github.com/golang/protobuf/proto"
	"github.com/profzone/imblock/global"
)

var _ interface {
	ProtocolSerializable
} = (*FindNodeAck)(nil)

func (msg *FindNodeAck) GetProtocolType() global.MessageType {
	return global.MESSAGE_TYPE__FIND_NODE_ACK
}

func (msg *FindNodeAck) EncodeFromSource() ([]byte, error) {
	return proto.Marshal(msg)
}

func (msg *FindNodeAck) DecodeFromSource(source []byte) error {
	return proto.Unmarshal(source, msg)
}