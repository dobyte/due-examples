package main

import (
	"github.com/dobyte/due-examples/cluster/service/grpc/internal/service/greeter/client"
	"github.com/dobyte/due-examples/cluster/service/grpc/internal/service/greeter/pb"
	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/registry/consul/v2"
	"github.com/dobyte/due/transport/grpc/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/codes"
	"github.com/dobyte/due/v2/log"
)

// 路由号
const hello = 1

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建用户定位器
	locator := redis.NewLocator()
	// 创建服务发现
	registry := consul.NewRegistry()
	// 创建RPC传输器
	transporter := grpc.NewTransporter()
	// 创建节点组件
	component := node.NewNode(
		node.WithLocator(locator),
		node.WithRegistry(registry),
		node.WithTransporter(transporter),
	)
	// 初始化应用
	initAPP(component.Proxy())
	// 添加节点组件
	container.Add(component)
	// 启动容器
	container.Serve()
}

// 初始化应用
func initAPP(proxy *node.Proxy) {
	proxy.Router().AddRouteHandler(hello, false, helloHandler)
}

// 请求
type helloReq struct {
	Name string `json:"name"`
}

// 响应
type helloRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 路由处理器
func helloHandler(ctx node.Context) {
	req := &helloReq{}
	res := &helloRes{}
	ctx.Defer(func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("response message failed: %v", err)
		}
	})

	if err := ctx.Parse(req); err != nil {
		log.Errorf("parse request message failed: %v", err)
		res.Code = codes.InternalError.Code()
		return
	}

	cli, err := client.NewClient(ctx.NewMeshClient)
	if err != nil {
		log.Errorf("create rpc client failed: %v", err)
		res.Code = codes.InternalError.Code()
		return
	}

	reply, err := cli.Hello(ctx.Context(), &pb.HelloArgs{Name: req.Name})
	if err != nil {
		log.Errorf("invoke rpc func failed: %v", err)
		res.Code = codes.Convert(err).Code()
		return
	}

	res.Code = codes.OK.Code()
	res.Message = reply.Message
}
