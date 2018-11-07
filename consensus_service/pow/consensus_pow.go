package pow

import "github.com/profzone/imblock/consensus_service"

var _ interface{
	consensus_service.ConsensusService
} = (*ConsensusPOW)(nil)

type ConsensusPOW struct {

}