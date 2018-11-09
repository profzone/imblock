package peers

import (
	"github.com/johnnyeven/libtools/courier"
)

var Router = courier.NewRouter(Peers{})

// 节点管理接口
type Peers struct {
	courier.EmptyOperator
}

func (req Peers) Path() string {
	return "/peers"
}
