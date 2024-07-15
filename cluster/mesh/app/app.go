package app

import (
	"due-examples/cluster/internal/service/rpcx/user"
	"github.com/dobyte/due/v2/cluster/mesh"
)

func Init(proxy *mesh.Proxy) {
	// 初始化用户服务
	user.NewServer(proxy).Init()
}
