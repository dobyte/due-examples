package main

import (
	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/registry/consul/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/codes"
	"github.com/dobyte/due/v2/log"
)

// 路由号
const (
	createRoom  = 1 // 创建房间
	destroyRoom = 2 // 销毁房间

)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建用户定位器
	locator := redis.NewLocator()
	// 创建服务发现
	registry := consul.NewRegistry()
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
	// 创建房间
	proxy.Router().AddRouteHandler(createRoom, false, createRoomHandler)
	// 销毁房间
	proxy.Router().AddRouteHandler(destroyRoom, true)
}

// 请求
type createRoomReq struct {
	RoomID string `json:"roomID"`
}

// 响应
type createRoomRes struct {
	Code int `json:"code"`
}

// 创建房间处理器
func createRoomHandler(ctx node.Context) {
	req := &createRoomReq{}
	res := &createRoomRes{}
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

	if _, err := ctx.Spawn(newRoomProcessor, node.WithActorID(req.RoomID), node.WithActorKind("room")); err != nil {
		log.Errorf("create actor failed: %v", err)
		res.Code = codes.Convert(err).Code()
		return
	}

	ctx.BindNode()

	res.Code = codes.OK.Code()
}

type roomProcessor struct {
	node.BaseProcessor
	actor *node.Actor
}

func newRoomProcessor(actor *node.Actor, args ...any) node.Processor {
	return &roomProcessor{actor: actor}
}

func (p *roomProcessor) Init() {
	p.actor.AddRouteHandler()
}
