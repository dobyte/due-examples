package main

import (
	"due-examples/cluster/client/app"
	"github.com/dobyte/due/eventbus/nats/v2"
	"github.com/dobyte/due/network/ws/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/client"
	"github.com/dobyte/due/v2/eventbus"
)

func main() {
	// 初始化事件总线
	eventbus.SetEventbus(nats.NewEventbus())
	// 创建容器
	container := due.NewContainer()
	// 创建客户端组件
	component := client.NewClient(
		client.WithClient(ws.NewClient()),
	)
	// 初始化应用
	app.Init(component.Proxy())
	// 添加客户端组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
