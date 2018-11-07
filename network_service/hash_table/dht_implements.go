package hash_table

import (
	"github.com/profzone/imblock/global"
	"github.com/profzone/imblock/global/protocols"
	"github.com/johnnyeven/terra/dht"
)

func Ping(node *dht.Node, t *dht.Transport) {

}

func Handshake(node *dht.Node, t *dht.Transport, target []byte) {
	message := &protocols.FindNode{
		Guid:       []byte(t.GetDHT().ID(string(target))),
		Version:    global.Config.Version,
		//Ip:         nil,
		//Port:       0,
		TargetGuid: target,
	}

	request := t.MakeRequest(node.ID, node.Addr, "", message)
	t.Request(request)
}
