package routers

import (
	"github.com/gin-gonic/gin"
	"zld-jy/action/users"
	"zld-jy/models"
)

func addUserRouters(rg *gin.RouterGroup) {
	rg.GET("/getUsers", func(context *gin.Context) {
		//处理逻辑
		var userId = context.Query("userId")
		userinfo := action_users.Instance.GetUserInfo(userId)
		models.OK(context, userinfo)
	})
	rg.POST("/register", func(context *gin.Context) {
		var registerUser models.RegisterUser
		context.BindJSON(&registerUser)
		action_users.Register(registerUser)
		models.OK(context, nil)
	})
}
