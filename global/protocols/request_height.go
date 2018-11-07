package protocols

import (
	"github.com/golang/protobuf/proto"
	"github.com/profzone/imblock/global"
)

var _ interface {
	ProtocolSerializable
} = (*RequestHeight)(nil)

func (msg *RequestHeight) GetProtocolType() global.MessageType {
	return global.MESSAGE_TYPE__REQUEST_HEIGHT
}

func (msg *RequestHeight) EncodeFromSource() ([]byte, error) {
	return proto.Marshal(msg)
}

func (msg *RequestHeight) DecodeFromSource(source []byte) error {
	return proto.Unmarshal(source, msg)
}