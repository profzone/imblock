package persistence_service

import (
	"github.com/profzone/imblock/core"
	"github.com/profzone/imblock/core/message_bus"
	"github.com/profzone/imblock/core/castle"
	"github.com/profzone/imblock/persistence_service/boltdb"
	"github.com/profzone/imblock/global"
	"github.com/sirupsen/logrus"
	"github.com/johnnyeven/libtools/courier"
	"github.com/profzone/imblock/core/constant"
	"errors"
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
		DB: boltdb.NewBoltDBBootstrap(),
	}

	return persistenceBoltDB
}

func (s *PersistenceBoltDBServiceBootstrap) Messages() []message_bus.MessageHandler {
	return []message_bus.MessageHandler{
		{
			Topic:  constant.TOPIC_PERSISTENCE_PUT,
			Runner: func(message message_bus.Message) (reply message_bus.Message, err error) {
				data, ok := message.Data.(map[string][]byte)
				if !ok {
					err = errors.New("[PersistenceBoltDBServiceBootstrap] HandlePut data should be map[string][]byte")
					return
				}

				err = s.DB.Put(data["bucket"], data["key"], data["value"])
				return
			},
		},
		{
			Topic:  constant.TOPIC_PERSISTENCE_GET,
			Runner: func(message message_bus.Message) (reply message_bus.Message, err error) {
				data, ok := message.Data.(map[string][]byte)
				if !ok {
					err = errors.New("[PersistenceBoltDBServiceBootstrap] HandleGet data should be map[string][]byte")
					return
				}

				result, err := s.DB.Get(data["bucket"], data["key"])
				if err != nil {
					return
				}
				reply.Data = result
				return
			},
		},
	}
}

func (s *PersistenceBoltDBServiceBootstrap) Protocols() []core.ProtocolHandler {
	return nil
}

func (s *PersistenceBoltDBServiceBootstrap) Routes() []*courier.Router {
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
