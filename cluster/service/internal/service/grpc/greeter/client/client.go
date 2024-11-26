package client

import (
	"due-examples/cluster/service/internal/service/grpc/greeter/pb"
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
