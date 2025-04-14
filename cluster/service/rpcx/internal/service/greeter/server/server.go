package server

import (
	"context"
	"due-examples/cluster/service/rpcx/internal/service/greeter/pb"
	"github.com/dobyte/due/v2/cluster/mesh"
)

const (
	service     = "greeter" // 用于客户端定位服务，例如discovery://greeter
	servicePath = "Greeter" // 服务路径要与pb中的服务路径保持一致
)

type Server struct {
	proxy *mesh.Proxy
}

var _ pb.GreeterAble = &Server{}

func NewServer(proxy *mesh.Proxy) *Server {
	return &Server{
		proxy: proxy,
	}
}

func (s *Server) Init() {
	s.proxy.AddServiceProvider(service, servicePath, s)
}

func (s *Server) Hello(ctx context.Context, args *pb.HelloArgs, reply *pb.HelloReply) (err error) {
	reply.Message = "Hello " + args.Name
	return
}
