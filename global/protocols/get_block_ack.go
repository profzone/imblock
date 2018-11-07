package protocols

import (
	"github.com/golang/protobuf/proto"
	"github.com/profzone/imblock/global"
	"github.com/profzone/imblock/core"
)

var _ interface {
	core.ProtocolSerializable
} = (*GetBlockAck)(nil)

func (msg *GetBlockAck) GetProtocolType() global.MessageType {
	return global.MESSAGE_TYPE__GET_BLOCK_ACK
}

func (msg *GetBlockAck) EncodeFromSource() ([]byte, error) {
	return proto.Marshal(msg)
}

func (msg *GetBlockAck) DecodeFromSource(source []byte) error {
	return proto.Unmarshal(source, msg)
}