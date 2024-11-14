package main

import (
	"fmt"
	"github.com/dobyte/due/component/http/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/codes"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/utils/xtime"
)

// @title API文档
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建HTTP组件
	component := http.NewServer()
	// 初始化监听
	initListen(component.Proxy())
	// 添加网格组件
	container.Add(component)
	// 启动容器
	container.Serve()
}

// 初始化监听
func initListen(proxy *http.Proxy) {
	// 路由器
	router := proxy.Router()
	// 注册路由
	router.Get("/greet", greetHandler)
}

// 请求
type greetReq struct {
	Message string `json:"message"`
}

// 响应
type greetRes struct {
	Message string `json:"message"`
}

// 测试接口
// @Summary 测试接口
// @Tags 测试
// @Schemes
// @Accept json
// @Produce json
// @Param request body greetReq true "请求参数"
// @Response 200 {object} http.Resp{Data=greetRes} "响应参数"
// @Router /greet [get]
func greetHandler(ctx http.Context) error {
	req := &greetReq{}

	if err := ctx.Bind().JSON(req); err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	log.Info(req.Message)

	return ctx.Success(&greetRes{
		Message: fmt.Sprintf("I'm server, and the current time is: %s", xtime.Now().Format(xtime.DateTime)),
	})
}
