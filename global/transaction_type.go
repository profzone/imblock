package global

//go:generate libtools gen enum TransactionType
// swagger:enum
type TransactionType uint8

// 交易类型
const (
	TRANSACTION_TYPE_UNKNOWN   TransactionType = iota
	TRANSACTION_TYPE__COINBASE  // 共识奖励交易
	TRANSACTION_TYPE__TRANSFER  // 普通转账交易
	TRANSACTION_TYPE__DATA      // 业务数据承载交易
)
