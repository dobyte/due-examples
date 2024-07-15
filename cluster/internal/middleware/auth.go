package middleware

import (
	"due-examples/cluster/internal/codes"
	"due-examples/cluster/internal/protocol/common"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/log"
)

func Auth(middleware *node.Middleware, ctx node.Context) {
	if ctx.UID() == 0 {
		if err := ctx.Response(&common.Res{Code: int32(codes.Unauthorized.Code())}); err != nil {
			log.Errorf("response message failed, err: %v", err)
		}
	} else {
		middleware.Next(ctx)
	}
}
