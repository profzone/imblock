package utxo

import (
	"github.com/profzone/imblock/core/model"
	"errors"
	"github.com/sirupsen/logrus"
	"encoding/hex"
)

type UTXOBootstrap struct {

}

func (b *UTXOBootstrap) CommitTransaction(transaction *model.Transaction) error {
	if transaction == nil {
		return errors.New("parameter <transaction> is nil")
	}

	// 保存ledger
	err := commitLedger(transaction.GetLedger())
	if err != nil {
		logrus.Errorf("[UTXOBootstrap] CommitTransaction commitLedger err: %v", err)
		rollbackErr := rollbackLedger(transaction.GetLedger())
		if rollbackErr != nil {
			logrus.Errorf("[UTXOBootstrap] CommitTransaction rollbackLedger err: %v", err)
		}
		logrus.Infof("[UTXOBootstrap] CommitTransaction rollbackLedger success")
		return err
	}

	// 保存交易
	err = utxoTransactionStorage.SaveTransaction(transaction)
	if err != nil {
		logrus.Errorf("[UTXOBootstrap] CommitTransaction SaveTransaction err: %v", err)
		rollbackErr := b.RollbackTransaction(transaction)
		if rollbackErr != nil {
			logrus.Errorf("[UTXOBootstrap] CommitTransaction RollbackTransaction err: %v", err)
		}
		logrus.Infof("[UTXOBootstrap] CommitTransaction RollbackTransaction success")
		return err
	}

	return nil
}

func commitLedger(ledger *model.Ledger) error {
	return nil
}

func rollbackLedger(ledger *model.Ledger) error {
	return nil
}

func (*UTXOBootstrap) RollbackTransaction(transaction *model.Transaction) error {
	if transaction == nil {
		return errors.New("parameter <transaction> is nil")
	}

	// 回滚ledger
	err := rollbackLedger(transaction.GetLedger())
	if err != nil {
		logrus.Errorf("[UTXOBootstrap] RollbackTransaction rollbackLedger err: %v", err)
		return err
	}

	// 回滚交易
	err = utxoTransactionStorage.DeleteTransaction(transaction.GetHash())
	if err != nil {
		logrus.Errorf("[UTXOBootstrap] RollbackTransaction DeleteTransaction err: %v", err)
		return err
	}

	return nil
}

func (*UTXOBootstrap) GetTransaction(hash []byte) *model.Transaction {
	if hash == nil {
		return nil
	}

	tx, err := utxoTransactionStorage.GetTransaction(hash)
	if err != nil {
		logrus.Errorf("[UTXOBootstrap] GetTransaction err: %v", err)
		return nil
	}

	return tx
}

func (b *UTXOBootstrap) VerifyTransactionLedger(transaction *model.Transaction) error {
	if transaction == nil {
		logrus.Debug("[UTXOBootstrap] VerifyTransactionLedger err: transaction is nil")
		return errors.New("transaction is nil")
	}

	ledger := transaction.GetLedger()
	if ledger == nil {
		logrus.Debug("[UTXOBootstrap] VerifyTransactionLedger err: ledger is nil")
		return errors.New("ledger is nil")
	}

	prevTrans := make(map[string]*model.Transaction)

	// 检查数据库中是否存在此交易的依赖交易
	from := ledger.GetFrom()
	for _, cell := range from {
		prevTranHash := cell.GetTranHash()
		prevTran := b.GetTransaction(prevTranHash)
		if prevTran == nil {
			logrus.Debugf("[UTXOBootstrap] VerifyTransactionLedger err: prevTran not found, currentCell: %s", cell)
			return errors.New("orphan transaction")
		}
		prevTrans[hex.EncodeToString(prevTran.GetHash())] = prevTran
	}

	// 验证交易签名


	return nil
}



