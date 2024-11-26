package app

import (
	"due-examples/cluster/actor/node/app/logic"
	"github.com/dobyte/due/v2/cluster/node"
)

func Init(proxy *node.Proxy) {
	logic.NewCore(proxy).Init()
}
