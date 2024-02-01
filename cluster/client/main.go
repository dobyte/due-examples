package main

import (
	"fmt"
	"github.com/dobyte/due/eventbus/nats/v2"
	"github.com/dobyte/due/network/ws/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/client"
	"github.com/dobyte/due/v2/eventbus"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/utils/xtime"
	"time"
)

const greet = 1

func main() {
	// 初始化事件总线
	eventbus.SetEventbus(nats.NewEventbus())
	// 创建容器
	container := due.NewContainer()
	// 创建客户端组件
	component := client.NewClient(
		client.WithClient(ws.NewClient()),
	)
	// 初始化监听
	initListen(component.Proxy())
	// 添加客户端组件
	container.Add(component)
	// 启动容器
	container.Serve()
}

// 初始化监听
func initListen(proxy *client.Proxy) {
	// 监听组件启动
	proxy.AddHookListener(cluster.Start, startHandler)
	// 监听连接建立
	proxy.AddEventListener(cluster.Connect, connectHandler)
	// 监听消息回复
	proxy.AddRouteHandler(greet, greetHandler)
}

// 组件启动处理器
func startHandler(proxy *client.Proxy) {
	if _, err := proxy.Dial(); err != nil {
		log.Errorf("gate connect failed: %v", err)
		return
	}
}

// 连接建立处理器
func connectHandler(conn *client.Conn) {
	doPushMessage(conn)
}

// 消息回复处理器
func greetHandler(ctx *client.Context) {
	res := &greetRes{}

	if err := ctx.Parse(res); err != nil {
		log.Errorf("invalid response message, err: %v", err)
		return
	}

	if res.Code != 0 {
		log.Errorf("node response failed, code: %d", res.Code)
		return
	}

	log.Info(res.Message)

	time.AfterFunc(time.Second, func() {
		doPushMessage(ctx.Conn())
	})
}

// 推送消息
func doPushMessage(conn *client.Conn) {
	err := conn.Push(&cluster.Message{
		Route: 1,
		Data: &greetReq{
			Message: fmt.Sprintf("I'm client, and the current time is: %s", xtime.Now().Format(xtime.DatetimeLayout)),
		},
	})
	if err != nil {
		log.Errorf("push message failed: %v", err)
	}
}

type greetReq struct {
	Message string `json:"message"`
}

type greetRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
