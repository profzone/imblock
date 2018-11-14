package utxo

import "github.com/profzone/imblock/core/model"

type UTXOBoostrap struct {

}

func (*UTXOBoostrap) CommitTransaction(transaction *model.Transaction) error {
	panic("implement me")
}

func (*UTXOBoostrap) RollbackTransaction(transaction *model.Transaction) error {
	panic("implement me")
}

func (*UTXOBoostrap) GetTransaction(hash []byte) {
	panic("implement me")
}

func (*UTXOBoostrap) VerifyTransactionPayload(transaction *model.Transaction) error {
	panic("implement me")
}



