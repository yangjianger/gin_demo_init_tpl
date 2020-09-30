package admin_router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//后台路由
func IncludeAdminRouters(r *gin.Engine) {
	r.GET("/admin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "admin",
		})
	})
}
