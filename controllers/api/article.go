package api

import (
	"blog_tpl/config"
	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
}

func NewArticleController() *ArticleController {
	return &ArticleController{}
}

func (this *ArticleController) Index(c *gin.Context) {

	ResponseSuccess(c, "article-module-index--"+config.Conf.MySQLConfig.Host)
}
