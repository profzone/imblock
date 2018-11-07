package protocols

import (
	"github.com/golang/protobuf/proto"
	"github.com/profzone/imblock/global"
)

var _ interface {
	ProtocolSerializable
} = (*GetBlock)(nil)

func (msg *GetBlock) GetProtocolType() global.MessageType {
	return global.MESSAGE_TYPE__GET_BLOCK
}

func (msg *GetBlock) EncodeFromSource() ([]byte, error) {
	return proto.Marshal(msg)
}

func (msg *GetBlock) DecodeFromSource(source []byte) error {
	return proto.Unmarshal(source, msg)
}