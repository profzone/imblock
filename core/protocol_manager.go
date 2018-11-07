package core

import (
	"github.com/sirupsen/logrus"
	"github.com/profzone/imblock/global"
	"github.com/johnnyeven/terra/dht"
)

type ProtocolRunner func(t *dht.Transport, msg *Protocol) error

type ProtocolHandler struct {
	Type   global.MessageType
	Runner ProtocolRunner
}

type ProtocolManager struct {
	protocolMap *dht.SyncedMap
}

var protocolManager *ProtocolManager

func GetProtocolManager() *ProtocolManager {
	if protocolManager == nil {
		protocolManager = &ProtocolManager{
			protocolMap: dht.NewSyncedMap(),
		}
	}

	return protocolManager
}

func (m *ProtocolManager) RegisterProtocol(handler ProtocolHandler) {
	if m.protocolMap.Has(handler.Type) {
		logrus.Panicf("[ProtocolManager] %s already registered", handler.Type.String())
	}

	m.protocolMap.Set(handler.Type, handler.Runner)
}

func (m *ProtocolManager) GetProtocolRunner(messageType global.MessageType) ProtocolRunner {
	v, ok := m.protocolMap.Get(messageType)
	if !ok {
		logrus.Panicf("[GetProtocolRunner] error: not found MessageType: %s", messageType.String())
	}
	return v.(ProtocolRunner)
}

