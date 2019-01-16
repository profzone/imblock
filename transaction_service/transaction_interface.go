package transaction_service

import "github.com/profzone/imblock/core/model"

type TransactionService interface {
	CommitTransaction(transaction *model.Transaction) error
	RollbackTransaction(transaction *model.Transaction) error
	GetTransaction(hash []byte) *model.Transaction
	VerifyTransactionLedger(transaction *model.Transaction) error
}