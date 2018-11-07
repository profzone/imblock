package protocols

import (
	"github.com/golang/protobuf/proto"
	"github.com/profzone/imblock/global"
	"github.com/profzone/imblock/core"
)

var _ interface {
	core.ProtocolSerializable
} = (*HeartbeatAck)(nil)

func (msg *HeartbeatAck) GetProtocolType() global.MessageType {
	return global.MESSAGE_TYPE__HEARTBEAT_ACK
}

func (msg *HeartbeatAck) EncodeFromSource() ([]byte, error) {
	return proto.Marshal(msg)
}

func (msg *HeartbeatAck) DecodeFromSource(source []byte) error {
	return proto.Unmarshal(source, msg)
}