package client

import (
	"github.com/dobyte/due-examples/cluster/service/grpc/internal/service/greeter/pb"
	"github.com/dobyte/due/v2/transport"
	"google.golang.org/grpc"
)

const target = "discovery://greeter"

func NewClient(fn transport.NewMeshClient) (pb.GreeterClient, error) {
	client, err := fn(target)
	if err != nil {
		return nil, err
	}

	return pb.NewGreeterClient(client.Client().(grpc.ClientConnInterface)), nil
}
