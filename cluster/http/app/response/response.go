package response

import (
	"github.com/dobyte/due/v2/codes"
	"github.com/dobyte/due/v2/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Res struct {
	Code int `json:"code"`
	Data any `json:"data,omitempty"`
}

const ExitPanic = "http response and exit"

// Fail 返回错误
func Fail(ctx *gin.Context, code *codes.Code) {
	Response(ctx, code)
}

// Success 返回成功
func Success(ctx *gin.Context, data ...any) {
	Response(ctx, codes.OK, data...)
}

// Response 响应消息
func Response(ctx *gin.Context, code *codes.Code, data ...any) {
	log.Debug(code.String())

	if len(data) > 0 {
		ctx.JSON(http.StatusOK, &Res{Code: code.Code(), Data: data[0]})
	} else {
		ctx.JSON(http.StatusOK, &Res{Code: code.Code()})
	}

	//panic(ExitPanic)
}
