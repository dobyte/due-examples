package client

import (
	"due-examples/cluster/service/rpcx/internal/service/greeter/pb"
	"github.com/dobyte/due/v2/transport"
	"github.com/smallnest/rpcx/client"
)

const target = "discovery://greeter"

func NewClient(fn transport.NewMeshClient) (*pb.GreeterOneClient, error) {
	c, err := fn(target)
	if err != nil {
		return nil, err
	}

	return pb.NewGreeterOneClient(c.Client().(*client.OneClient)), nil
}
