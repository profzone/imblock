package hash_table

import (
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
	"time"
	"github.com/profzone/imblock/core"
	"github.com/johnnyeven/terra/dht"
)

func DHTPacketHandler(table *dht.DistributedHashTable, packet dht.Packet) {
	msg := unpackMessageFromPackage(packet.Data)
	runner := core.GetProtocolManager().GetProtocolRunner(msg.Header)

	logrus.Debug("[DHTPacketHandler] Handle message [MsgHeader=", msg.Header.String(), ", MsgID=", msg.ProtocolID, "] started")

	err := runner(table.GetTransport(), msg)
	if err != nil {
		logrus.Errorf("[DHTPacketHandler] Handle message err: %v", err)
	}

	logrus.Debug("[DHTPacketHandler] Handle message [MsgHeader=", msg.Header.String(), ", MsgID=", msg.ProtocolID, "] ended")
}

func NewNodeHandler(peerID []byte, node *dht.Node) {
	logrus.Infof("new node recorded, id: %x, ip: %s, port: %d", []byte(node.ID.RawString()), node.Addr.IP.String(), node.Addr.Port)

	//TODO PeerManager count limit
	if peerManager.Has(peerID) {
		return
	}

	remote := net.JoinHostPort(node.Addr.IP.String(), strconv.FormatUint(uint64(node.Addr.Port), 10))
	conn, err := net.DialTimeout("tcp", remote, 10*time.Second)
	if err != nil {
		logrus.Errorf("[NewNodeHandler] net.DialTimeout error: %v", err)
	}

	p := NewPeerWithConnection(peerID, node, conn.(*net.TCPConn))
	peerManager.Set(p)
	logrus.Infof("new peer connect id: %x, address: %s", p.Guid, conn.RemoteAddr().String())

	p.Handshake()
}
