package main

import (
	"due-examples/cluster/service/internal/service/rpcx/greeter/server"
	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/registry/consul/v2"
	"github.com/dobyte/due/transport/rpcx/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/mesh"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建用户定位器
	locator := redis.NewLocator()
	// 创建服务发现
	registry := consul.NewRegistry()
	// 创建RPC传输器
	transporter := rpcx.NewTransporter()
	// 创建网格组件
	component := mesh.NewMesh(
		mesh.WithLocator(locator),
		mesh.WithRegistry(registry),
		mesh.WithTransporter(transporter),
	)
	// 初始化应用
	initRPCXServer(component.Proxy())
	// 添加网格组件
	container.Add(component)
	// 启动容器
	container.Serve()
}

// 初始化应用
func initRPCXServer(proxy *mesh.Proxy) {
	server.NewServer(proxy).Init()
}
