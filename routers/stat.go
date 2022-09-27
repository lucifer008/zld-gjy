package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addStatsRouters(cg *gin.RouterGroup) {
	cg.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "stats")
	})
}
