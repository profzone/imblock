package protocols

import (
	"github.com/golang/protobuf/proto"
	"github.com/profzone/imblock/global"
	"github.com/profzone/imblock/core"
)

var _ interface {
	core.ProtocolSerializable
} = (*BlocksHash)(nil)

func (msg *BlocksHash) GetProtocolType() global.MessageType {
	return global.MESSAGE_TYPE__BLOCKS_HASH
}

func (msg *BlocksHash) EncodeFromSource() ([]byte, error) {
	return proto.Marshal(msg)
}

func (msg *BlocksHash) DecodeFromSource(source []byte) error {
	return proto.Unmarshal(source, msg)
}