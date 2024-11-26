package logic

import (
	"due-examples/cluster/actor/node/app/route"
	"github.com/dobyte/due/v2/cluster/node"
)

type Core struct {
	proxy *node.Proxy
}

func NewCore(proxy *node.Proxy) *Core {
	return &Core{proxy: proxy}
}

func (c *Core) Init() {
	c.proxy.AddRouteHandler(route.CreateRoom, false, c.createRoom)
}
