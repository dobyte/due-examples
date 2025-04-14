package main

import (
	"fmt"
	"github.com/dobyte/due/network/tcp/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/client"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/utils/xtime"
	"time"
)

// 路由号
const greet = 1

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建客户端组件
	component := client.NewClient(
		client.WithClient(tcp.NewClient()),
	)
	// 初始化监听
	initListen(component.Proxy())
	// 添加客户端组件
	container.Add(component)
	// 启动容器
	container.Serve()
}

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
		log.Errorf("connect server failed: %v", err)
		return
	}
}

// 连接建立处理器
func connectHandler(conn *client.Conn) {
	pushMessage(conn)
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
		pushMessage(ctx.Conn())
	})
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

// 推送消息
func pushMessage(conn *client.Conn) {
	err := conn.Push(&cluster.Message{
		Route: 1,
		Data: &greetReq{
			Message: fmt.Sprintf("I'm tcp client, and the current time is: %s", xtime.Now().Format(xtime.DateTime)),
		},
	})
	if err != nil {
		log.Errorf("push message failed: %v", err)
	}
}
