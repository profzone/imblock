package protocols

import "github.com/profzone/imblock/global"

type ProtocolSerializable interface {
	GetProtocolType() global.MessageType
	DecodeFromSource([]byte) error
	EncodeFromSource() ([]byte, error)
}
