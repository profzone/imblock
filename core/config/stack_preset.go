package config

import (
	"github.com/profzone/imblock/network_service"
	"github.com/profzone/imblock/core/castle"
)

var PresetStack = map[string][]castle.ServiceConstructor{
	"fnode": {
		network_service.NewNetworkDHTServiceBootstrap,
		//NewHeartbeatService,
		//NewBlockChainService,
	},
}
