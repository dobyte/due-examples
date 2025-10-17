package main

import (
	"fmt"

	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/registry/etcd/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/codes"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/utils/xtime"
)

// 路由号
const greet = 1

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建用户定位器
	locator := redis.NewLocator()
	// 创建服务发现
	registry := etcd.NewRegistry()
	// 创建节点组件
	component := node.NewNode(
		node.WithLocator(locator),
		node.WithRegistry(registry),
	)
	// 初始化应用
	initApp(component.Proxy())
	// 添加节点组件
	container.Add(component)
	// 启动容器
	container.Serve()
}

// 初始化应用
func initApp(proxy *node.Proxy) {
	proxy.Router().Group(func(group *node.RouterGroup) {
		group.Middleware(Auth)

		group.AddRouteHandler(greet, greetHandler)
	})
}

func Auth(middleware *node.Middleware, ctx node.Context) {
	middleware.Next(ctx)
}

// 请求
type greetReq struct {
	Message string `json:"message"`
}

// 响应
type greetRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 路由处理器
func greetHandler(ctx node.Context) {
	ctx.Task(func(ctx node.Context) {
		req := &greetReq{}
		res := &greetRes{}
		defer func() {
			if err := ctx.Response(res); err != nil {
				log.Errorf("response message failed: %v", err)
			}
		}()

		if err := ctx.Parse(req); err != nil {
			log.Errorf("parse request message failed: %v", err)
			res.Code = codes.InternalError.Code()
			return
		}

		log.Info(req.Message)

		res.Code = codes.OK.Code()
		res.Message = fmt.Sprintf("I'm ws server, and the current time is: %s", xtime.Now().Format(xtime.DateTime))
	})
}
