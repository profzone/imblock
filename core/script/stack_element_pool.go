package script

import (
	"math/big"
	"sync"
)

const poolLimit = 256
const poolsCap = 10

type stackElementPool struct {
	pool *Stack
}

func newStackElementPool() *stackElementPool {
	return &stackElementPool{newStack()}
}

func (p *stackElementPool) get() *big.Int {
	if p.pool.len() > 0 {
		return p.pool.pop().SetUint64(0)
	}

	return new(big.Int)
}

func (p *stackElementPool) put(values ...*big.Int) {
	if p.pool.len() >= poolLimit {
		return
	}

	p.pool.pushN(values...)
}

var stackElementPools = &stackElementPoolPool{
	pools: make([]*stackElementPool, 0, poolsCap),
}

type stackElementPoolPool struct {
	pools []*stackElementPool
	lock  sync.Mutex
}

func (pp *stackElementPoolPool) get() *stackElementPool {
	pp.lock.Lock()
	defer pp.lock.Unlock()

	if len(pp.pools) > 0 {
		p := pp.pools[len(pp.pools)-1]
		pp.pools = pp.pools[:len(pp.pools)-1]
		return p
	}

	return newStackElementPool()
}

func (pp *stackElementPoolPool) put(p *stackElementPool) {
	pp.lock.Lock()
	defer pp.lock.Unlock()

	if len(pp.pools) < cap(pp.pools) {
		pp.pools = append(pp.pools, p)
	}
}
