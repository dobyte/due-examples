package logic

import (
	"due-examples/cluster/actor/node/app/route"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/codes"
	"github.com/dobyte/due/v2/log"
)

type Core struct {
	proxy *node.Proxy
	mgr   *RoomMgr
}

func NewCore(proxy *node.Proxy) *Core {
	return &Core{
		proxy: proxy,
		mgr:   NewRoomMgr(proxy),
	}
}

func (c *Core) Init() {
	// 创建房间
	c.proxy.AddRouteHandler(route.CreateRoom, false, c.createRoom)
	// 搜索房间
	c.proxy.AddRouteHandler(route.SearchRoom, false, c.searchRoom)
	// 进入房间
	c.proxy.AddRouteHandler(route.EnterRoom, false, c.enterRoom)
	// 离开聊天室
	c.proxy.AddRouteHandler(route.LeaveRoom, false, c.leaveRoom)
}

// 创建房间
func (c *Core) createRoom(ctx node.Context) {
	req := &createRoomReq{}
	res := &createRoomRes{}
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

	room, err := c.mgr.createRoom(ctx.UID(), req.Name)
	if err != nil {
		log.Errorf("fetch table list failed: %v", err)
		res.Code = codes.Convert(err).Code()
		return
	}

	res.Code = codes.OK.Code()
	res.Data = &createRoomResData{Room: room.makeRoomInfo()}
}

// 搜索房间
func (c *Core) searchRoom(ctx node.Context) {
	req := &createRoomReq{}
	res := &createRoomRes{}
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

	room, err := c.mgr.createRoom(ctx.UID(), req.Name)
	if err != nil {
		res.Code = codes.Convert(err).Code()
		return
	}

	res.Code = codes.OK.Code()
	res.Data = &createRoomResData{Room: room.makeRoomInfo()}
}

// 进入聊天室
func (c *Core) enterRoom(ctx node.Context) {
	req := &enterRoomReq{}
	res := &enterRoomRes{}
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

	if err := c.mgr.enterRoom(ctx.UID(), req.ID); err != nil {
		res.Code = codes.Convert(err).Code()
		return
	}

	res.Code = codes.OK.Code()
}

func (c *Core) leaveRoom(ctx node.Context) {

}
