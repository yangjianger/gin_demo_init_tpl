package home_router

import (
	"github.com/gin-gonic/gin"
)

//后台路由
func IncludeHomeRouters(r *gin.Engine) {
	apiRouter := r.Group("/api")

	//注册文章路由
	includeArticleRouter(apiRouter)

}
