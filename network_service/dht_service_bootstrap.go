package network_service

import (
	"net"
	"github.com/sirupsen/logrus"
	"github.com/profzone/imblock/global"
	"github.com/profzone/imblock/core/castle"
	"github.com/profzone/imblock/core/message_bus"
	"github.com/profzone/imblock/network_service/hash_table"
	"github.com/profzone/imblock/core"
	"github.com/johnnyeven/terra/dht"
	"github.com/johnnyeven/libtools/courier"
)

var networkDHT *NetworkDHTServiceBootstrap

var _ interface {
	castle.Service
} = (*NetworkDHTServiceBootstrap)(nil)

type NetworkDHTServiceBootstrap struct {
	listener   net.Listener
	service    NetworkService
	quitSignal chan struct{}
}

func GetChainDHTConfig() *dht.Config {
	return &dht.Config{
		BucketExpiredAfter:   global.Config.BucketExpiredAfter,
		NodeExpiredAfter:     global.Config.NodeExpriedAfter,
		CheckBucketPeriod:    global.Config.CheckBucketPeriod,
		MaxTransactionCursor: global.Config.MaxTransactionCursor,
		MaxNodes:             global.Config.MaxNodes,
		K:                    global.Config.K,
		BucketSize:           global.Config.BucketSize,
		RefreshNodeCount:     global.Config.RefreshNodeCount,
		Network:              global.Config.Network,
		LocalAddr:            global.Config.LocalAddr.String(),
		SeedNodes:            global.Config.SeedNodes,
		TransportConstructor: hash_table.NewProtobufTransport,
		NewNodeHandler:       hash_table.NewNodeHandler,
		Handler:              hash_table.DHTPacketHandler,
		HandshakeFunc:        hash_table.Handshake,
		PingFunc:             hash_table.Ping,
	}
}

func NewNetworkDHTServiceBootstrap() castle.Service {
	if networkDHT != nil {
		return networkDHT
	}

	table := dht.NewDHT(GetChainDHTConfig())

	networkDHT = &NetworkDHTServiceBootstrap{
		service:    table,
		quitSignal: make(chan struct{}),
	}

	return networkDHT
}

func (s *NetworkDHTServiceBootstrap) listen() {
	var err error
	var localAddr = global.Config.LocalAddr.String()
	s.listener, err = net.Listen("tcp", localAddr)
	if err != nil {
		logrus.Panicf("[NetworkDHTServiceBootstrap] net.Listen error: %v", err.Error())
	}
	logrus.Debugf("[NetworkDHTServiceBootstrap] created and listened at: %s ...", localAddr)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			logrus.Fatalf("[NetworkDHTServiceBootstrap] listener.Accept error: %v", err)
		}

		go s.handleConnection(conn)
	}
}

func (s *NetworkDHTServiceBootstrap) handleConnection(conn net.Conn) {
	p := hash_table.NewPeerWithConnection(nil, nil, conn.(*net.TCPConn))
	logrus.Infof("[NetworkDHTServiceBootstrap] new peer connected id: %x, address: %s", p.Guid, conn.RemoteAddr().String())
}

func (s *NetworkDHTServiceBootstrap) Messages() []message_bus.MessageHandler {
	return nil
}

func (s *NetworkDHTServiceBootstrap) Protocols() []core.ProtocolHandler {
	return []core.ProtocolHandler{
		{
			Type:   global.MESSAGE_TYPE__HELLO_TCP,
			Runner: hash_table.RunHelloTCP,
		},
		{
			Type:   global.MESSAGE_TYPE__FIND_NODE,
			Runner: hash_table.RunFindNode,
		},
		{
			Type:   global.MESSAGE_TYPE__FIND_NODE_ACK,
			Runner: hash_table.RunFindNodeAck,
		},
		{
			Type:   global.MESSAGE_TYPE__HEARTBEAT,
			Runner: hash_table.RunHeartbeat,
		},
		{
			Type:   global.MESSAGE_TYPE__HEARTBEAT_ACK,
			Runner: hash_table.RunHeartbeatAck,
		},
	}
}

func (s *NetworkDHTServiceBootstrap) Routes() []*courier.Router {
	return nil
}

func (s *NetworkDHTServiceBootstrap) Start() error {

	go s.listen()
	go s.service.Run()

	logrus.Info("[NetworkDHTServiceBootstrap] started.")
Run:
	for {
		select {
		case <-s.quitSignal:
			break Run
		}
	}

	close(s.quitSignal)
	return nil
}

func (s *NetworkDHTServiceBootstrap) Stop() error {

	s.quitSignal <- struct{}{}
	s.service.Close()

	logrus.Info("[NetworkDHTServiceBootstrap] stopped.")
	return nil
}
