package logic

import (
	"due-examples/cluster/hall/app/route"
	"due-examples/cluster/internal/codes"
	"due-examples/cluster/internal/protocol/hall"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/client"
	"github.com/dobyte/due/v2/log"
)

type core struct {
	proxy *client.Proxy
}

func NewCode(proxy *client.Proxy) *core {
	return &core{proxy: proxy}
}

func (c *core) Init() {
	// 监听组件启动
	c.proxy.AddHookListener(cluster.Start, c.startHandler)
	// 监听连接建立
	c.proxy.AddEventListener(cluster.Connect, c.connectHandler)
	// 监听注册消息回复
	c.proxy.AddRouteHandler(route.Register, c.registerHandler)
	// 监听登录消息回复
	c.proxy.AddRouteHandler(route.Login, c.loginHandler)
	// 监听拉取资料消息回复
	c.proxy.AddRouteHandler(route.FetchProfile, c.fetchProfileHandler)
}

// 组件启动处理器
func (c *core) startHandler(proxy *client.Proxy) {
	if _, err := proxy.Dial(); err != nil {
		log.Errorf("gate connect failed: %v", err)
		return
	}
}

// 连接建立处理器
func (c *core) connectHandler(conn *client.Conn) {
	log.Info("gate connect success")

	if err := conn.Push(&cluster.Message{
		Route: route.Register,
		Data: &hall.RegisterReq{
			Account:  "fuxiao",
			Password: "123456",
			Nickname: "月半拂晓",
		},
	}); err != nil {
		log.Errorf("push register message failed: %v", err)
	}
}

// 监听注册消息回复
func (c *core) registerHandler(ctx *client.Context) {
	res := &hall.RegisterRes{}

	if err := ctx.Parse(res); err != nil {
		log.Errorf("invalid response message, err: %v", err)
		return
	}

	switch int(res.Code) {
	case codes.OK.Code():
		log.Info("account register success")

		c.pushLoginMessage(ctx)
	case codes.AccountExists.Code():
		log.Info("account exits")

		c.pushLoginMessage(ctx)
	default:
		log.Errorf("node response failed, code: %d", res.Code)

	}
}

func (c *core) pushLoginMessage(ctx *client.Context) {
	if err := ctx.Conn().Push(&cluster.Message{
		Route: route.Login,
		Data: &hall.LoginReq{
			Account:  "fuxiao",
			Password: "123456",
		},
	}); err != nil {
		log.Errorf("push login message failed: %v", err)
	}
}

// 监听登录消息回复
func (c *core) loginHandler(ctx *client.Context) {
	res := &hall.LoginRes{}

	if err := ctx.Parse(res); err != nil {
		log.Errorf("invalid response message, err: %v", err)
		return
	}

	if res.Code != 0 {
		log.Errorf("node response failed, code: %d", res.Code)
		return
	}

	log.Info("account login success")

	if err := ctx.Conn().Push(&cluster.Message{
		Route: route.FetchProfile,
	}); err != nil {
		log.Errorf("push fetch profile message failed: %v", err)
	}
}

// 监听拉取资料消息回复
func (c *core) fetchProfileHandler(ctx *client.Context) {
	res := &hall.FetchProfileRes{}

	if err := ctx.Parse(res); err != nil {
		log.Errorf("invalid response message, err: %v", err)
		return
	}

	if res.Code != 0 {
		log.Errorf("node response failed, code: %d", res.Code)
		return
	}

	log.Info("fetch profile success")
	log.Infof("UID: %d", res.Data.UID)
	log.Infof("Account: %s", res.Data.Account)
	log.Infof("Nickname: %s", res.Data.Nickname)
}
