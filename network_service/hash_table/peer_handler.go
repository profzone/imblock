package hash_table

import (
	"github.com/sirupsen/logrus"
	"github.com/profzone/imblock/core"
	"github.com/johnnyeven/terra/dht"
)

func PeerPacketHandler(p *Peer, packet dht.Packet) {
	msg := unpackMessageFromPackage(packet.Data)
	runner := core.GetProtocolManager().GetProtocolRunner(msg.Header)

	logrus.Debug("[PeerPacketHandler] Handle message [MsgHeader=", msg.Header.String(), ", MsgID=", msg.ProtocolID, "] started")

	err := runner(p.transport, msg)
	if err != nil {
		logrus.Errorf("[PeerPacketHandler] Handle message err: %v", err)
	}

	tranID := msg.ProtocolID
	tran := p.transport.Get(tranID, msg.RemoteAddr)
	if tran != nil {
		defer func() {
			tran.ResponseChannel <- struct{}{}
		}()
	}

	logrus.Debug("[PeerPacketHandler] Handle message [MsgHeader=", msg.Header.String(), ", MsgID=", msg.ProtocolID, "] ended")
}
