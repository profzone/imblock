package consensus_service

import (
	"github.com/profzone/imblock/core/castle"
	"github.com/profzone/imblock/core/message_bus"
	"github.com/profzone/imblock/global/protocols"
)

var _ interface{
	castle.Service
} = (*ConsensusPOWServiceBootstrap)(nil)

type ConsensusPOWServiceBootstrap struct {
	service ConsensusService
}

func (s *ConsensusPOWServiceBootstrap) Messages() []message_bus.MessageHandler {
	return nil
}

func (s *ConsensusPOWServiceBootstrap) Protocols() []protocols.ProtocolHandler {
	return nil
}

func (s *ConsensusPOWServiceBootstrap) Start() error {
	return nil
}

func (s *ConsensusPOWServiceBootstrap) Stop() error {
	return nil
}
