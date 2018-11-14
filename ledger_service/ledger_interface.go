package ledger_service

import "github.com/profzone/imblock/core/model"

type LedgerService interface {
	CommitTransaction(transaction *model.Transaction) error
	RollbackTransaction(transaction *model.Transaction) error
	GetTransaction(hash []byte)
	VerifyTransactionPayload(transaction *model.Transaction) error
}