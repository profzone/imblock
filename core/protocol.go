package core

import (
	"net"
	"encoding/gob"
	"github.com/profzone/imblock/global"
)

type Protocol struct {
	Header     global.MessageType
	PeerID     []byte
	ProtocolID int64
	Payload    []byte
	RemoteAddr net.Addr
}

func init() {
	gob.Register(&net.UDPAddr{})
	gob.Register(&net.TCPAddr{})
}

type ProtocolSerializable interface {
	GetProtocolType() global.MessageType
	DecodeFromSource([]byte) error
	EncodeFromSource() ([]byte, error)
}