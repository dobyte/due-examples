package main

import (
	"context"
	"due-examples/cluster/service/grpc/internal/service/greeter/client"
	pb2 "due-examples/cluster/service/grpc/internal/service/greeter/pb"
	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/registry/consul/v2"
	"github.com/dobyte/due/transport/grpc/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/codes"
	"github.com/dobyte/due/v2/log"
	"time"
)

// 路由号
const hello = 1

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建用户定位器
	locator := redis.NewLocator()
	// 创建服务发现
	registry := consul.NewRegistry()
	// 创建RPC传输器
	transporter := grpc.NewTransporter()
	// 创建节点组件
	component := node.NewNode(
		node.WithLocator(locator),
		node.WithRegistry(registry),
		node.WithTransporter(transporter),
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
	proxy.Router().AddRouteHandler(hello, false, helloHandler)

	NewServer(proxy).Init()
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

// 路由处理器
func helloHandler(ctx node.Context) {
	req := &helloReq{}
	res := &helloRes{}
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

	cli, err := client.NewClient(ctx.Proxy().NewMeshClient)
	if err != nil {
		log.Errorf("create rpc client failed: %v", err)
		res.Code = codes.InternalError.Code()
		return
	}

	reply, err := cli.Hello(ctx.Context(), &pb2.HelloArgs{Name: req.Name})
	if err != nil {
		log.Errorf("invoke rpc func failed: %v", err)
		res.Code = codes.Convert(err).Code()
		return
	}

	ctx.AfterFunc(time.Second, func() {
		log.Debugf("after exec success")
	})

	res.Code = codes.OK.Code()
	res.Message = reply.Message
}

// ///////////////////grpc server
type Server struct {
	pb2.UnimplementedGreeterServer
	proxy *node.Proxy
}

var _ pb2.GreeterServer = &Server{}

func NewServer(proxy *node.Proxy) *Server {
	return &Server{
		proxy: proxy,
	}
}

func (s *Server) Init() {
	s.proxy.AddServiceProvider("greeter", &pb2.Greeter_ServiceDesc, s)
}

func (s *Server) Hello(_ context.Context, args *pb2.HelloArgs) (*pb2.HelloReply, error) {
	return &pb2.HelloReply{Message: "Hello " + args.Name}, nil
}
