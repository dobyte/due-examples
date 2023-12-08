package main

import (
	"fmt"
	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/registry/consul/v2"
	"github.com/dobyte/due/transport/rpcx/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/codes"
	"github.com/dobyte/due/v2/log"
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
	// 创建网关组件
	component := node.NewNode(
		node.WithLocator(locator),
		node.WithRegistry(registry),
		node.WithTransporter(transporter),
	)
	// 注册路由
	component.Proxy().Router().AddRouteHandler(1, false, greetHandler)
	// 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}

type greetReq struct {
	Name string `json:"name"`
}

type greetRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func greetHandler(ctx *node.Context) {
	req := &greetReq{}
	res := &greetRes{}
	defer func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("response message failed: %v", err)
		}
	}()

	if err := ctx.Request.Parse(req); err != nil {
		log.Errorf("parse request message failed: %v", err)
		res.Code = codes.InternalError.Code()
		return
	}

	res.Code = codes.OK.Code()
	res.Message = fmt.Sprintf("hello %s~~", req.Name)
}
