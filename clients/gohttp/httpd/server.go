package httpd

import (
	"github.com/opentdp/go-helper/httpd"

	"wechat-rest/args"
	"wechat-rest/httpd/midware"
	"wechat-rest/httpd/wcfrest"
)

// @title Wechat Rest API
// @version v0.4.6
// @description 基于 WeChatFerry RPC 实现的电脑版微信 REST API，使用 Go 语言编写，无第三方运行时依赖。基于 HTTP 提供操作接口，轻松对接任意编程语言。
// @contact.name WeChatRest
// @contact.url https://github.com/opentdp/wechat-rest
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api

func Server() {

	httpd.Engine(args.Debug)

	// Api 守卫
	api := httpd.Group("/api")
	api.Use(midware.OutputHandle, midware.ApiGuard)

	// Wcf 路由
	wcfrest.Route(api)

	// Swagger 守卫
	httpd.Use(midware.SwagGuard)

	// 前端文件路由
	httpd.StaticEmbed("/", "public", args.Efs)

	// 启动 HTTP 服务
	httpd.Server(args.Httpd.Address)

}
