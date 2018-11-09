package persistence_service

import (
	"github.com/profzone/imblock/core"
	"github.com/profzone/imblock/core/message_bus"
	"github.com/profzone/imblock/core/castle"
	"github.com/profzone/imblock/persistence_service/boltdb"
	"github.com/profzone/imblock/global"
	"github.com/sirupsen/logrus"
)

var persistenceBoltDB *PersistenceBoltDBServiceBootstrap

type PersistenceBoltDBServiceBootstrap struct {
	DB PersistenceService
}

func NewPersistenceBoltDBServiceBootstrap() castle.Service {
	if persistenceBoltDB != nil {
		return persistenceBoltDB
	}

	persistenceBoltDB = &PersistenceBoltDBServiceBootstrap{
		DB: boltdb.NewBoltDBBootstrap(global.BlockBucketIdentity),
	}

	return persistenceBoltDB
}

func (s *PersistenceBoltDBServiceBootstrap) Messages() []message_bus.MessageHandler {
	return nil
}

func (s *PersistenceBoltDBServiceBootstrap) Protocols() []core.ProtocolHandler {
	return nil
}

func (s *PersistenceBoltDBServiceBootstrap) Start() error {
	err := s.DB.Open(global.GetBlockFilePath(), nil)
	if err == nil {
		logrus.Infof("[PersistenceBoltDBServiceBootstrap] started, database file path: %s", global.GetBlockFilePath())
	} else {
		logrus.Errorf("[PersistenceBoltDBServiceBootstrap] start failed, err: %v", err)
	}

	return err
}

func (s *PersistenceBoltDBServiceBootstrap) Stop() error {
	err := s.DB.Close()
	if err == nil {
		logrus.Info("[PersistenceBoltDBServiceBootstrap] stopped.")
	} else {
		logrus.Errorf("[PersistenceBoltDBServiceBootstrap] stop failed, err: %v", err)
	}

	return err
}
