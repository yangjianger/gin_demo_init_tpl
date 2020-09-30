package api

import (
	"blog_tpl/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
	controllers.BaseController
}

func (this *BaseController) Index(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "base-module-index",
	})

}
