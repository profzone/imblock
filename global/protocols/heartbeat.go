package protocols

import (
	"github.com/golang/protobuf/proto"
	"github.com/profzone/imblock/global"
)

var _ interface {
	ProtocolSerializable
} = (*Heartbeat)(nil)

func (msg *Heartbeat) GetProtocolType() global.MessageType {
	return global.MESSAGE_TYPE__HEARTBEAT
}

func (msg *Heartbeat) EncodeFromSource() ([]byte, error) {
	return proto.Marshal(msg)
}

func (msg *Heartbeat) DecodeFromSource(source []byte) error {
	return proto.Unmarshal(source, msg)
}
