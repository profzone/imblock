package core

import (
	"github.com/sirupsen/logrus"
	"github.com/johnnyeven/terra/dht"
)

func PeerPacketHandler(p *Peer, packet dht.Packet) {
	msg := UnpackMessageFromPackage(packet.Data)
	runner := protocolManager.GetProtocolRunner(msg.Header)

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
