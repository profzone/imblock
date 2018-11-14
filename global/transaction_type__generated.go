package global

import (
	"bytes"
	"encoding"
	"errors"

	github_com_johnnyeven_libtools_courier_enumeration "github.com/johnnyeven/libtools/courier/enumeration"
)

var InvalidTransactionType = errors.New("invalid TransactionType")

func init() {
	github_com_johnnyeven_libtools_courier_enumeration.RegisterEnums("TransactionType", map[string]string{
		"COINBASE": "共识奖励交易",
		"DATA":     "业务数据承载交易",
		"TRANSFER": "普通转账交易",
	})
}

func ParseTransactionTypeFromString(s string) (TransactionType, error) {
	switch s {
	case "":
		return TRANSACTION_TYPE_UNKNOWN, nil
	case "COINBASE":
		return TRANSACTION_TYPE__COINBASE, nil
	case "DATA":
		return TRANSACTION_TYPE__DATA, nil
	case "TRANSFER":
		return TRANSACTION_TYPE__TRANSFER, nil
	}
	return TRANSACTION_TYPE_UNKNOWN, InvalidTransactionType
}

func ParseTransactionTypeFromLabelString(s string) (TransactionType, error) {
	switch s {
	case "":
		return TRANSACTION_TYPE_UNKNOWN, nil
	case "共识奖励交易":
		return TRANSACTION_TYPE__COINBASE, nil
	case "业务数据承载交易":
		return TRANSACTION_TYPE__DATA, nil
	case "普通转账交易":
		return TRANSACTION_TYPE__TRANSFER, nil
	}
	return TRANSACTION_TYPE_UNKNOWN, InvalidTransactionType
}

func (TransactionType) EnumType() string {
	return "TransactionType"
}

func (TransactionType) Enums() map[int][]string {
	return map[int][]string{
		int(TRANSACTION_TYPE__COINBASE): {"COINBASE", "共识奖励交易"},
		int(TRANSACTION_TYPE__DATA):     {"DATA", "业务数据承载交易"},
		int(TRANSACTION_TYPE__TRANSFER): {"TRANSFER", "普通转账交易"},
	}
}
func (v TransactionType) String() string {
	switch v {
	case TRANSACTION_TYPE_UNKNOWN:
		return ""
	case TRANSACTION_TYPE__COINBASE:
		return "COINBASE"
	case TRANSACTION_TYPE__DATA:
		return "DATA"
	case TRANSACTION_TYPE__TRANSFER:
		return "TRANSFER"
	}
	return "UNKNOWN"
}

func (v TransactionType) Label() string {
	switch v {
	case TRANSACTION_TYPE_UNKNOWN:
		return ""
	case TRANSACTION_TYPE__COINBASE:
		return "共识奖励交易"
	case TRANSACTION_TYPE__DATA:
		return "业务数据承载交易"
	case TRANSACTION_TYPE__TRANSFER:
		return "普通转账交易"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*TransactionType)(nil)

func (v TransactionType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidTransactionType
	}
	return []byte(str), nil
}

func (v *TransactionType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseTransactionTypeFromString(string(bytes.ToUpper(data)))
	return
}
