package logic

import (
	"due-examples/cluster/actor/app/route"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/node"
)

type core struct {
	proxy *node.Proxy
}

func NewCore(proxy *node.Proxy) *core {
	return &core{
		proxy: proxy,
	}
}

func (c *core) Init() {
	c.proxy.Router().Group(func(group *node.RouterGroup) {
		// 注册中间件
		group.Middleware(middleware.State, middleware.Auth)
		// 创建房间
		group.AddRouteHandler(route.CreateRoom, false, c.createRoom)
		// 解散房间
		group.AddRouteHandler(route.DismissRoom, false, c.dismissRoom)
		// 创建牌桌
		group.AddRouteHandler(route.CreateTable, true, c.createTable)
		// 退出游戏
		group.AddRouteHandler(route.DismissTable, true, c.exitGame)
		// 玩家出牌
		group.AddRouteHandler(route.PlayCards, true, c.playCards)
		// 玩家选牌
		group.AddRouteHandler(route.SelectCards, true, c.selectCards)
	})

	// 断线重连
	c.proxy.AddEventHandler(cluster.Reconnect, c.reconnect)
	// 断开连接
	c.proxy.AddEventHandler(cluster.Disconnect, c.disconnect)
}

// 创建房间
func (c *core) createRoom(ctx node.Context) {
	pid := ctx.Actor().Spawn(newRoom)
}

// 解散房间
func (c *core) dismissRoom(ctx node.Context) {
	pid := ctx.Actor().Load()
	if pid == nil {
		return
	}

}

// 创建牌桌
func (c *core) createTable(ctx node.Context) {

}
