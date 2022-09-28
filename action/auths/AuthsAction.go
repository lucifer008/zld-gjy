package auths

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"zld-jy/models"
)

type AuthsHandler struct {
}

// Login
//@Summary 登录接口
// @Schemes
// @Description 登录
// @Tags ops 认证
// @Param user body models.Users true "登录参数"
// @Accept json
// @Produce json
// @Success 200
// @Router /auths/login [post]
func (ah AuthsHandler) Login(c *gin.Context) {
	log.Printf("请求成功")
	var users models.Users
	c.IndentedJSON(http.StatusOK, users)
	log.Printf(">>>>>解析后的参数:", users)
}
