package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"zld-gly/models"
)

func addAuthsRouters(rg *gin.RouterGroup) {
	rg.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "auths")
	})
	rg.POST("/login", func(context *gin.Context) {
		var user models.Users
		if err := context.BindJSON(&user); err != nil {
			log.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", user.Username)
			log.Default().Fatal(err)
		}
		log.Printf(">>>>>>>>>>>请求成功! 用户名:", user.Username, "密码:", user.Password)
	})
}
