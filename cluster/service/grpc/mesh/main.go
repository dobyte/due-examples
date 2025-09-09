package main

import (
	"context"

	"github.com/dobyte/due-examples/cluster/service/grpc/internal/service/greeter/server"
	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/registry/consul/v2"
	"github.com/dobyte/due/transport/grpc/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/mesh"
	"github.com/dobyte/due/v2/log"
	ggrpc "google.golang.org/grpc"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建用户定位器
	locator := redis.NewLocator()
	// 创建服务发现
	registry := consul.NewRegistry()
	// 创建RPC传输器
	transporter := grpc.NewTransporter(grpc.WithServerOptions(ggrpc.ChainUnaryInterceptor(loggerInterceptor, authInterceptor)))
	// 创建网格组件
	component := mesh.NewMesh(
		mesh.WithLocator(locator),
		mesh.WithRegistry(registry),
		mesh.WithTransporter(transporter),
	)
	// 初始化应用
	initAPP(component.Proxy())
	// 添加网格组件
	container.Add(component)
	// 启动容器
	container.Serve()
}

// 日志拦截器
func loggerInterceptor(ctx context.Context, req any, info *ggrpc.UnaryServerInfo, handler ggrpc.UnaryHandler) (any, error) {
	// 打印请求日志
	log.Debugf("logger interceptor received request: %v", req)
	// 调用下一个拦截器
	return handler(ctx, req)
}

// 授权拦截器
func authInterceptor(ctx context.Context, req any, info *ggrpc.UnaryServerInfo, handler ggrpc.UnaryHandler) (any, error) {
	// 打印请求日志
	log.Debugf("auth interceptor received request: %v", req)
	// 调用下一个拦截器
	return handler(ctx, req)
}

// 初始化应用
func initAPP(proxy *mesh.Proxy) {
	server.NewServer(proxy).Init()
}
