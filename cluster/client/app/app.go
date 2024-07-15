package app

import (
	"due-examples/cluster/client/app/logic"
	"github.com/dobyte/due/v2/cluster/client"
)

func Init(proxy *client.Proxy) {
	logic.NewCode(proxy).Init()
}
