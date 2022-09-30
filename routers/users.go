package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"zld-jy/action/users"
)

func addUserRouters(rg *gin.RouterGroup) {
	rg.GET("/", func(context *gin.Context) {
		//处理逻辑
		action_users.Instance.GetUserInfo(context)
		context.JSON(http.StatusOK, "users:"+time.Now().String())
	})
}
