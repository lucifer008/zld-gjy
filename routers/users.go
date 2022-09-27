package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addUserRouters(rg *gin.RouterGroup) {
	rg.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "users")
	})
}
