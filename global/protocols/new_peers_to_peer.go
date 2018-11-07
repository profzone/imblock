package protocols

import (
	"github.com/golang/protobuf/proto"
	"github.com/profzone/imblock/global"
	"github.com/profzone/imblock/core"
)

var _ interface {
	core.ProtocolSerializable
} = (*NewPeersToPeer)(nil)

func (msg *NewPeersToPeer) GetProtocolType() global.MessageType {
	return global.MESSAGE_TYPE__NEW_PEERS_TO_PEER
}

func (msg *NewPeersToPeer) EncodeFromSource() ([]byte, error) {
	return proto.Marshal(msg)
}

func (msg *NewPeersToPeer) DecodeFromSource(source []byte) error {
	return proto.Unmarshal(source, msg)
}
