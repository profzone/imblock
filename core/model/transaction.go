package model

import (
	"github.com/profzone/imblock/global"
	"fmt"
	"time"
)

type Transaction struct {
	// 交易类型
	txType global.TransactionType
	// 交易账本 - utxo模型
	ledger *Ledger
	// 交易业务数据承载
	txData TransactionData
	// 交易哈希
	hash []byte
	// 交易时间戳
	timestamp time.Time
	// 交易签名
	signature []byte
}

func (tx *Transaction) GetHash() []byte {
	return tx.hash
}

func (tx *Transaction) GetLedger() *Ledger {
	return tx.ledger
}

func (tx *Transaction) GetSignature() []byte {
	return tx.signature
}

func (tx *Transaction) Serialize() ([]byte, error) {
	panic("implement me")
}

func (tx *Transaction) Deserialize([]byte) error {
	panic("implement me")
}

func (tx *Transaction) String() string {
	return fmt.Sprintf("Transaction{ type=%s, ledger=%s, txData=%s, timestamp=%s, hash=%s, signature=%s }",
		tx.txType.String(),
		tx.ledger.String(),
		tx.txData.String(),
		tx.timestamp.String(),
		fmt.Sprintf("0x%x", tx.hash),
		fmt.Sprintf("0x%x", tx.signature))
}
