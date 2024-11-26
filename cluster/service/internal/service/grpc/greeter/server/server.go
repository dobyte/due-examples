package server

import (
	"context"
	"due-examples/cluster/service/internal/service/grpc/greeter/pb"
	"github.com/dobyte/due/v2/cluster/mesh"
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
	return &pb.HelloReply{Message: "Hello " + args.Name}, nil
}
