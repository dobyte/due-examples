package main

import (
	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/log/zap/v2"
	"github.com/dobyte/due/network/ws/v2"
	"github.com/dobyte/due/registry/consul/v2"
	"github.com/dobyte/due/transport/rpcx/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/gate"
	"github.com/dobyte/due/v2/component/pprof"
	"github.com/dobyte/due/v2/log"
)

func main() {
	// 设置日志
	log.SetLogger(zap.NewLogger(zap.WithCallerSkip(2)))
	// 创建容器
	container := due.NewContainer()
	// 创建服务器
	server := ws.NewServer()
	// 创建用户定位器
	locator := redis.NewLocator()
	// 创建服务发现
	registry := consul.NewRegistry()
	// 创建RPC传输器
	transporter := rpcx.NewTransporter()
	// 创建网关组件
	component1 := gate.NewGate(
		gate.WithServer(server),
		gate.WithLocator(locator),
		gate.WithRegistry(registry),
		gate.WithTransporter(transporter),
	)
	// 创建性能分析组件
	component2 := pprof.NewPProf()
	// 添加网关组件
	container.Add(component1, component2)
	// 启动容器
	container.Serve()
}
