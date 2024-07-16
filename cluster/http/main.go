package main

import (
	"due-examples/cluster/http/app"
	"github.com/dobyte/due/component/http/v2"
	"github.com/dobyte/due/registry/consul/v2"
	"github.com/dobyte/due/transport/grpc/v2"
	"github.com/dobyte/due/v2"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建服务发现
	registry := consul.NewRegistry()
	// 创建RPC传输器
	transporter := grpc.NewTransporter()
	// 创建HTTP组件
	component := http.NewHttp(
		http.WithAddr(":3993"),
		http.WithRegistry(registry),
		http.WithTransporter(transporter),
	)
	// 初始化应用
	app.Init(component.Proxy())
	// 添加网格组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
