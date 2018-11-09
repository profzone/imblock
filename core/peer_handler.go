package core

import (
	"github.com/sirupsen/logrus"
	"github.com/johnnyeven/terra/dht"
)

func peerPacketHandler(p *Peer, packet dht.Packet) {
	msg := UnpackMessageFromPackage(packet.Data)
	runner := protocolManager.GetProtocolRunner(msg.Header)

	logrus.Debug("[peerPacketHandler] Handle message [MsgHeader=", msg.Header.String(), ", MsgID=", msg.ProtocolID, "] started")

	err := runner(p.transport, msg)
	if err != nil {
		logrus.Errorf("[peerPacketHandler] Handle message err: %v", err)
	}

	tranID := msg.ProtocolID
	tran := p.transport.Get(tranID, msg.RemoteAddr)
	if tran != nil {
		defer func() {
			tran.ResponseChannel <- struct{}{}
		}()
	}

	logrus.Debug("[peerPacketHandler] Handle message [MsgHeader=", msg.Header.String(), ", MsgID=", msg.ProtocolID, "] ended")
}
