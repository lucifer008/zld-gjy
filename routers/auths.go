package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addAuthsRouters(rg *gin.RouterGroup) {
	rg.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "auths")
	})
}
