package routers

import (
	"github.com/gin-gonic/gin"
	"zld-jy/action"
	"zld-jy/models"
)

func addUserRouters(rg *gin.RouterGroup) {
	rg.GET("/getUsers", func(context *gin.Context) {
		//处理逻辑
		var userId = context.Query("userId")
		userinfo := action.UserActionInstance.GetUserInfo(userId)
		models.OK(context, userinfo)
	})
	rg.POST("/register", func(context *gin.Context) {
		var registerUser models.RegisterUser
		context.BindJSON(&registerUser)
		error := action.UserActionInstance.Register(registerUser)
		if error != nil {
			panic(error)
		}
		models.OK(context, nil)
	})
}
