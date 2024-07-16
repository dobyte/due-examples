package app

import (
	"context"
	"due-examples/cluster/http/app/response"
	"due-examples/cluster/internal/service/grpc/user"
	"due-examples/cluster/internal/service/grpc/user/pb"
	"github.com/dobyte/due/component/http/v2"
	"github.com/dobyte/due/v2/codes"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/utils/xconv"
	"github.com/gin-gonic/gin"
)

type Res struct {
}

func Init(proxy *http.Proxy) {
	base := proxy.Engine().Group("/")

	base.GET("/great", func(ctx *gin.Context) {
		uid, ok := ctx.GetQuery("uid")
		if !ok {
			response.Fail(ctx, codes.InvalidArgument)
		}

		client, err := user.NewClient(proxy.NewMeshClient)
		if err != nil {
			log.Errorf("create mesh client failed: %v", err)
			response.Fail(ctx, codes.InternalError)
			return
		}

		reply, err := client.FetchProfile(context.Background(), &pb.FetchProfileRequest{
			UID: xconv.Int64(uid),
		})
		if err != nil {
			response.Fail(ctx, codes.Convert(err))
			return
		}

		response.Success(ctx, map[string]string{
			"account":  reply.Account,
			"nickname": reply.Nickname,
		})
	})
}
