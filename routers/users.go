package routers

import (
	"github.com/gin-gonic/gin"
	"zld-jy/action/users"
	"zld-jy/models"
)

func addUserRouters(rg *gin.RouterGroup) {
	rg.GET("/getUsers", func(context *gin.Context) {
		//处理逻辑
		var userId = context.Param("userId")
		userinfo := action_users.Instance.GetUserInfo(userId)
		models.OK(context, userinfo)
	})
}
