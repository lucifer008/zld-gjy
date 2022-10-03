package auths

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"zld-jy/models"
)

var Instance *AuthsAction

func init() {
	Instance = &AuthsAction{}
}

type AuthsAction struct {
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
func (ah AuthsAction) Login(c *gin.Context) {
	var users models.Users
	if err := c.BindJSON(&users); err != nil {
		c.JSONP(http.StatusOK, gin.H{"status": "错误:" + err.Error()})
		return
	}

	log.Printf(">>>>>解析后的参数:", users)
}
