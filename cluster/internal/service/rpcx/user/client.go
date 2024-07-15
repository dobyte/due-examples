package user

import (
	"due-examples/cluster/internal/service/rpcx/user/pb"
	"github.com/dobyte/due/v2/transport"
	"github.com/smallnest/rpcx/client"
)

const target = "discovery://user"

func NewClient(fn transport.NewMeshClient) (*pb.UserOneClient, error) {
	c, err := fn(target)
	if err != nil {
		return nil, err
	}

	return pb.NewUserOneClient(c.Client().(*client.OneClient)), nil
}
