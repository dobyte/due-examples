package server

import (
	"context"
	pb2 "due-examples/cluster/service/grpc/internal/service/greeter/pb"
	"github.com/dobyte/due/v2/cluster/mesh"
)

type Server struct {
	pb2.UnimplementedGreeterServer
	proxy *mesh.Proxy
}

var _ pb2.GreeterServer = &Server{}

func NewServer(proxy *mesh.Proxy) *Server {
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
