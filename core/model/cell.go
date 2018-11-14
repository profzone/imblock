package model

type Cells struct {
	inputs  []Cell
	outputs []Cell
}

func (cs Cells) String() string {
	panic("implement me")
}

type Cell struct {
	amount   uint64
	lockTime int64
	script   []byte
}

func (Cell) String() string {
	panic("implement me")
}

