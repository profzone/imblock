package protocols

import (
	"github.com/golang/protobuf/proto"
	"github.com/profzone/imblock/global"
)

var _ interface {
	ProtocolSerializable
} = (*HelloTCP)(nil)

func (msg *HelloTCP) GetProtocolType() global.MessageType {
	return global.MESSAGE_TYPE__HELLO_TCP
}

func (msg *HelloTCP) EncodeFromSource() ([]byte, error) {
	return proto.Marshal(msg)
}

func (msg *HelloTCP) DecodeFromSource(source []byte) error {
	return proto.Unmarshal(source, msg)
}
