package utxo

import (
	"github.com/profzone/imblock/core/model"
	"github.com/profzone/imblock/core/message_bus"
	"github.com/profzone/imblock/core"
	"github.com/profzone/imblock/core/constant"
	"errors"
)

var utxoTransactionStorage utxoTransactionStorageBootstrap

func init() {
	utxoTransactionStorage = utxoTransactionStorageBootstrap{
		messageBus: core.GetMessageManager(),
	}
}

type utxoTransactionStorageBootstrap struct {
	messageBus *message_bus.MessageBus
}

func (b utxoTransactionStorageBootstrap) SaveTransaction(transaction *model.Transaction) error {
	txBytes, err := transaction.Serialize()
	if err != nil {
		return err
	}

	msg := message_bus.Message{
		Topic: constant.TOPIC_PERSISTENCE_PUT,
		Data: map[string][]byte{
			"bucket": []byte(DB_BUCKET_TX),
			"key":    transaction.GetHash(),
			"value":  txBytes,
		},
	}
	_, err = b.messageBus.Publish(msg)
	if err != nil {
		return err
	}
	return nil
}

func (b utxoTransactionStorageBootstrap) GetTransaction(hash []byte) (*model.Transaction, error) {

	msg := message_bus.Message{
		Topic: constant.TOPIC_PERSISTENCE_GET,
		Data: map[string][]byte{
			"bucket": []byte(DB_BUCKET_TX),
			"key":    hash,
		},
	}
	replies, err := b.messageBus.Publish(msg)
	if err != nil {
		return nil, err
	}
	txBytes, ok := replies[0].Data.([]byte)
	if !ok {
		return nil, errors.New("[utxoTransactionStorageBootstrap] HandlePut reply data should be []byte")
	}
	tx := &model.Transaction{}
	err = tx.Deserialize(txBytes)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (b utxoTransactionStorageBootstrap) DeleteTransaction(hash []byte) error {
	return nil
}
