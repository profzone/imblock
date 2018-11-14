package model

import "fmt"

type TransactionData interface {
	fmt.Stringer
	Size() int
	Serialize() ([]byte, error)
	Deserialize([]byte) error
}
