package home_router

import (
	"blog_tpl/controllers/api"
	"github.com/gin-gonic/gin"
)

//文章路由
func includeArticleRouter(r *gin.RouterGroup) {
	articleRouter := r.Group("/article")

	articleController := api.NewArticleController()

	//首页
	articleRouter.GET("/index", articleController.Index)
}
