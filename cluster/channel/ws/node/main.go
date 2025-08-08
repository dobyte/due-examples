package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/registry/etcd/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/utils/xtime"
)

// 路由号
const greetNotify = 1

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
	proxy.AddEventHandler(cluster.Connect, connectHandler)

	go publishMessage(proxy)
}

func connectHandler(ctx node.Context) {
	channel := fmt.Sprintf("channel-%d", ctx.CID()%3)

	ctx.Subscribe(channel)
}

// 发布频道消息
func publishMessage(proxy *node.Proxy) {
	for {
		time.Sleep(time.Second)

		channel := fmt.Sprintf("channel-%d", time.Now().Unix()%3)

		proxy.Publish(context.Background(), &cluster.PublishArgs{
			Channel: channel,
			Message: &cluster.Message{
				Route: greetNotify,
				Data:  &greetNotice{Message: fmt.Sprintf("channel: %s time: %s", channel, xtime.Now().Format(xtime.DateTime))},
			},
		})
	}
}

type greetNotice struct {
	Message string `json:"message"`
}
