package hash_table

import (
	"github.com/johnnyeven/terra/dht"
	"github.com/profzone/imblock/core"
	"github.com/profzone/imblock/global/protocols"
	"github.com/sirupsen/logrus"
	"strings"
	"bytes"
	"github.com/profzone/imblock/global"
	"errors"
)

func RunHelloTCP(t *dht.Transport, msg *core.Protocol) error {

	payload := &protocols.HelloTCP{}
	err := payload.DecodeFromSource(msg.Payload)
	if err != nil {
		return err
	}

	node, err := dht.NewNodeFromCompactInfo(string(payload.Node), "udp4")
	if err != nil {
		return err
	}

	p := t.GetClient().(*ChainProtobufClient).GetPeer()
	p.Node = node
	p.Guid = payload.Guid

	core.GetPeerManager().Set(p)

	logrus.Infof("peer handshake, peerID: %x", p.Guid)

	return nil
}

func RunFindNode(t *dht.Transport, msg *core.Protocol) error {

	payload := &protocols.FindNode{}
	err := payload.DecodeFromSource(msg.Payload)
	if err != nil {
		return err
	}

	targetID := dht.NewIdentityFromBytes(payload.TargetGuid)

	var nodes string
	node, _ := t.GetDHT().GetRoutingTable().GetNodeBucketByID(targetID)
	if node != nil {
		nodes = node.CompactNodeInfo()
	} else {
		nodes = strings.Join(t.GetDHT().GetRoutingTable().GetNeighborCompactInfos(targetID, t.GetDHT().K), "")
	}

	message := &protocols.FindNodeAck{
		Guid:    []byte(t.GetDHT().ID(targetID.RawString())),
		Version: global.Config.Version,
		//Ip:      network.P2P.AnnouncedAddr.IP,
		//Port:    uint32(network.P2P.AnnouncedAddr.Port),
		Nodes: []byte(nodes),
	}

	request := t.MakeResponse(nil, msg.RemoteAddr, msg.ProtocolID, message)
	t.GetClient().(*ProtobufClient).Send(request)

	if bytes.Compare(msg.PeerID, global.Config.Guid) == 0 {
		return nil
	}

	n, _ := dht.NewNode(string(payload.Guid), msg.RemoteAddr.Network(), msg.RemoteAddr.String())
	t.GetDHT().GetRoutingTable().Insert(n)

	return nil
}

func RunFindNodeAck(t *dht.Transport, msg *core.Protocol) error {

	payload := &protocols.FindNodeAck{}
	err := payload.DecodeFromSource(msg.Payload)
	if err != nil {
		return err
	}

	tranID := msg.ProtocolID
	tran := t.GetDHT().GetTransport().Get(tranID, msg.RemoteAddr)
	if tran == nil {
		return errors.New("error trans")
	}

	defer func() {
		tran.ResponseChannel <- struct{}{}
	}()

	guid := payload.Guid

	//if tran.ClientID.(*dht.Identity) != nil && tran.ClientID.(*dht.Identity).RawString() != string(guid) {
	//	t.GetDHT().GetRoutingTable().RemoveByAddr(msg.RemoteAddr.String())
	//	return nil
	//}

	node, err := dht.NewNode(string(guid), msg.RemoteAddr.Network(), msg.RemoteAddr.String())
	if err != nil {
		return err
	}

	if err := findOrContinueRequestTarget(t.GetDHT(), payload.Guid, payload); err != nil {
		return err
	}

	if bytes.Compare(msg.PeerID, global.Config.Guid) == 0 {
		return nil
	}

	if t.GetDHT().GetRoutingTable().Insert(node) {
		if t.GetDHT().NewNodeHandler != nil {
			t.GetDHT().NewNodeHandler(msg.PeerID, node)
		}
	}

	return nil
}

func findOrContinueRequestTarget(table *dht.DistributedHashTable, targetID []byte, data *protocols.FindNodeAck) error {

	nodes := string(data.Nodes)
	if len(nodes)%26 != 0 {
		return errors.New("the length of nodes should can be divided by 26")
	}

	hasNew, found := false, false
	for i := 0; i < len(nodes)/26; i++ {
		node, _ := dht.NewNodeFromCompactInfo(string(nodes[i*26:(i+1)*26]), table.Network)
		if bytes.Compare(node.ID.RawData(), targetID) == 0 {
			found = true
		}

		if table.GetRoutingTable().Insert(node) {
			hasNew = true
		}
	}
	if found || !hasNew {
		return nil
	}

	id := dht.NewIdentityFromBytes(targetID)
	for _, node := range table.GetRoutingTable().GetNeighbors(id, table.K) {
		Handshake(node, table.GetTransport(), targetID)
	}

	return nil
}
