package model

type Ledger struct {
	from []Cell
	to   []Cell
}

func (l Ledger) GetFrom() []Cell {
	return l.from
}

func (l *Ledger) GetTo() []Cell {
	return l.to
}

func (l Ledger) String() string {
	panic("implement me")
}

type Cell struct {
	tranHash []byte
	amount   uint64
	lockTime int64
	script   []byte
}

func (c Cell) GetTranHash() []byte {
	return c.tranHash
}

func (Cell) String() string {
	panic("implement me")
}
