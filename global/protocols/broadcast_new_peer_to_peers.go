package protocols

import (
	"github.com/golang/protobuf/proto"
	"github.com/profzone/imblock/global"
)

var _ interface {
	ProtocolSerializable
} = (*BroadcastNewPeerToPeers)(nil)

func (msg *BroadcastNewPeerToPeers) GetProtocolType() global.MessageType {
	return global.MESSAGE_TYPE__BROADCAST_NEW_PEER_TO_PEERS
}

func (msg *BroadcastNewPeerToPeers) EncodeFromSource() ([]byte, error) {
	return proto.Marshal(msg)
}

func (msg *BroadcastNewPeerToPeers) DecodeFromSource(source []byte) error {
	return proto.Unmarshal(source, msg)
}