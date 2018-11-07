package protocols

import (
	"github.com/golang/protobuf/proto"
	"github.com/profzone/imblock/global"
	"github.com/profzone/imblock/core"
)

var _ interface {
	core.ProtocolSerializable
} = (*NewTransaction)(nil)

func (msg *NewTransaction) GetProtocolType() global.MessageType {
	return global.MESSAGE_TYPE__NEW_TRANSACTION
}

func (msg *NewTransaction) EncodeFromSource() ([]byte, error) {
	return proto.Marshal(msg)
}

func (msg *NewTransaction) DecodeFromSource(source []byte) error {
	return proto.Unmarshal(source, msg)
}
