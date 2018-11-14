package model

import (
	"github.com/profzone/imblock/global"
	"fmt"
	"time"
)

type Transaction struct {
	// 交易类型
	txType global.TransactionType
	// 交易utxo模型
	cells *Cells
	// 交易业务数据承载
	txData TransactionData
	// 交易哈希
	hash []byte
	// 交易时间戳
	timestamp time.Time
	// 交易签名
	signature []byte
}

func (tx *Transaction) String() string {
	return fmt.Sprintf("Transaction{ type=%s, cells=%s, txData=%s, timestamp=%s, hash=%s, signature=%s }",
		tx.txType.String(),
		tx.cells.String(),
		tx.txData.String(),
		tx.timestamp.String(),
		fmt.Sprintf("0x%x", tx.hash),
		fmt.Sprintf("0x%x", tx.signature))
}
