package main

import (
	"time"

	"github.com/dobyte/due/network/ws/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/client"
	"github.com/dobyte/due/v2/log"
)

// 路由号
const hello = 1

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建客户端组件
	component := client.NewClient(
		client.WithClient(ws.NewClient()),
	)
	// 初始化应用
	initAPP(component.Proxy())
	// 添加客户端组件
	container.Add(component)
	// 启动容器
	container.Serve()
}

// 初始化应用
func initAPP(proxy *client.Proxy) {
	// 监听组件启动
	proxy.AddHookListener(cluster.Start, startHandler)
	// 监听连接建立
	proxy.AddEventListener(cluster.Connect, connectHandler)
	// 监听消息回复
	proxy.AddRouteHandler(hello, helloHandler)
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
func helloHandler(ctx *client.Context) {
	res := &helloRes{}

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
type helloReq struct {
	Name string `json:"name"`
}

// 响应
type helloRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 推送消息
func pushMessage(conn *client.Conn) {
	if err := conn.Push(&cluster.Message{
		Route: hello,
		Data:  &helloReq{Name: "client"},
	}); err != nil {
		log.Errorf("push message failed: %v", err)
	}
}
