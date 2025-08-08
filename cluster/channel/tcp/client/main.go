package main

import (
	"github.com/dobyte/due/network/tcp/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/client"
	"github.com/dobyte/due/v2/log"
)

// 路由号
const greetNotify = 1

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建客户端组件
	component := client.NewClient(
		client.WithClient(tcp.NewClient()),
	)
	// 初始化监听
	initApp(component.Proxy())
	// 添加客户端组件
	container.Add(component)
	// 启动容器
	container.Serve()
}

func initApp(proxy *client.Proxy) {
	// 监听组件启动
	proxy.AddHookListener(cluster.Start, startHandler)
	// 监听消息回复
	proxy.AddRouteHandler(greetNotify, greetNoticeHandler)
}

// 组件启动处理器
func startHandler(proxy *client.Proxy) {
	for range 10 {
		if _, err := proxy.Dial(); err != nil {
			log.Errorf("connect server failed: %v", err)
			return
		}
	}
}

// 消息通知处理器
func greetNoticeHandler(ctx *client.Context) {
	notice := &greetNotice{}

	if err := ctx.Parse(notice); err != nil {
		log.Errorf("invalid notice message, err: %v", err)
		return
	}

	log.Info(notice.Message)
}

type greetNotice struct {
	Message string `json:"message"`
}
