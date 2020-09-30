package routers

import (
	"blog_tpl/middlewares"
	"blog_tpl/routers/admin_router"
	"blog_tpl/routers/home_router"
	"github.com/gin-gonic/gin"
)

//配置路由信息
func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middlewares.GinLogger(), middlewares.GinRecovery(false))

	//后台路由
	admin_router.IncludeAdminRouters(r)

	//前台路由
	home_router.IncludeHomeRouters(r)

	return r
}
