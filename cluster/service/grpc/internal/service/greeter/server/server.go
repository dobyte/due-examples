package server

import (
	"context"

	"github.com/dobyte/due-examples/cluster/service/grpc/internal/service/greeter/pb"
	"github.com/dobyte/due/v2/cluster/mesh"
	"github.com/dobyte/due/v2/log"
)

type Server struct {
	pb.UnimplementedGreeterServer
	proxy *mesh.Proxy
}

var _ pb.GreeterServer = &Server{}

func NewServer(proxy *mesh.Proxy) *Server {
	return &Server{
		proxy: proxy,
	}
}

func (s *Server) Init() {
	s.proxy.AddServiceProvider("greeter", &pb.Greeter_ServiceDesc, s)
}

func (s *Server) Hello(_ context.Context, args *pb.HelloArgs) (*pb.HelloReply, error) {
	log.Debugf("server handler received request: %v", args)

	return &pb.HelloReply{Message: "Hello " + args.Name}, nil
}
