package logic

import (
	"due-examples/cluster/hall/app/route"
	"due-examples/cluster/internal/middleware"
	"due-examples/cluster/internal/protocol/hall"
	"due-examples/cluster/internal/service/rpcx/user"
	"due-examples/cluster/internal/service/rpcx/user/pb"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/codes"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/task"
)

type core struct {
	proxy *node.Proxy
}

func NewCore(proxy *node.Proxy) *core {
	return &core{
		proxy: proxy,
	}
}

func (c *core) Init() {
	c.proxy.Router().Group(func(group *node.RouterGroup) {
		// 注册账号
		group.AddRouteHandler(route.Register, false, c.register)
		// 登录账号
		group.AddRouteHandler(route.Login, false, c.login)
		// 注册中间件
		group.Middleware(middleware.Auth)
		// 拉取信息
		group.AddRouteHandler(route.FetchProfile, false, c.fetchProfile)
	})
}

// 注册账号
// 由于单个节点上的路由消息处理采用的单线程，如果要进行存在I/O操作，最好使用task.AddTask来进行处理
func (c *core) register(ctx node.Context) {
	ctx = ctx.Clone()

	task.AddTask(func() {
		req := &hall.RegisterReq{}
		res := &hall.RegisterRes{}
		defer func() {
			if err := ctx.Response(res); err != nil {
				log.Errorf("response message failed: %v", err)
			}
		}()

		if err := ctx.Parse(req); err != nil {
			log.Errorf("parse request message failed: %v", err)
			res.Code = int32(codes.InternalError.Code())
			return
		}

		client, err := user.NewClient(c.proxy.NewMeshClient)
		if err != nil {
			log.Errorf("create mesh client failed: %v", err)
			res.Code = int32(codes.InternalError.Code())
			return
		}

		_, err = client.Register(ctx.Context(), &pb.RegisterRequest{
			Account:  req.Account,
			Password: req.Password,
			Nickname: req.Nickname,
		})
		if err != nil {
			res.Code = int32(codes.Convert(err).Code())
			return
		}

		res.Code = int32(codes.OK.Code())
	})
}

// 登录账号
func (c *core) login(ctx node.Context) {
	ctx = ctx.Clone()

	task.AddTask(func() {
		req := &hall.LoginReq{}
		res := &hall.LoginRes{}
		defer func() {
			if err := ctx.Response(res); err != nil {
				log.Errorf("response message failed: %v", err)
			}
		}()

		if err := ctx.Parse(req); err != nil {
			log.Errorf("parse request message failed: %v", err)
			res.Code = int32(codes.InternalError.Code())
			return
		}

		client, err := user.NewClient(c.proxy.NewMeshClient)
		if err != nil {
			log.Errorf("create mesh client failed: %v", err)
			res.Code = int32(codes.InternalError.Code())
			return
		}

		reply, err := client.Login(ctx.Context(), &pb.LoginRequest{
			Account:  req.Account,
			Password: req.Password,
		})
		if err != nil {
			res.Code = int32(codes.Convert(err).Code())
			return
		}

		// 登录成功后，将网关连接与用户ID进行绑定，在后续通过该连接到达node节点的消息都会携带上此用户ID
		if err = ctx.BindGate(reply.UID); err != nil {
			log.Errorf("bind gate failed: %v", err)
			res.Code = int32(codes.InternalError.Code())
			return
		}

		res.Code = int32(codes.OK.Code())
	})
}

// 拉取信息
func (c *core) fetchProfile(ctx node.Context) {
	ctx = ctx.Clone()

	task.AddTask(func() {
		res := &hall.FetchProfileRes{}
		defer func() {
			if err := ctx.Response(res); err != nil {
				log.Errorf("response message failed: %v", err)
			}
		}()

		client, err := user.NewClient(c.proxy.NewMeshClient)
		if err != nil {
			log.Errorf("create mesh client failed: %v", err)
			res.Code = int32(codes.InternalError.Code())
			return
		}

		reply, err := client.FetchProfile(ctx.Context(), &pb.FetchProfileRequest{
			UID: ctx.UID(),
		})
		if err != nil {
			res.Code = int32(codes.Convert(err).Code())
			return
		}

		res.Code = int32(codes.OK.Code())
		res.Data = &hall.Profile{
			UID:      ctx.UID(),
			Account:  reply.Account,
			Nickname: reply.Nickname,
		}
	})
}
