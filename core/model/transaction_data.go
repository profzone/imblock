package model

import "fmt"

type TransactionData interface {
	fmt.Stringer
	Serializable
	Size() int
}
