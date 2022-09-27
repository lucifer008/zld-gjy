package auths

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"zld-gly/models"
)

type AuthsHandler struct {
}

func (ah AuthsHandler) Login(c *gin.Context) {
	log.Printf("请求成功")
	var users models.Users
	c.IndentedJSON(http.StatusOK, users)
	log.Printf(">>>>>解析后的参数:", users)
}
