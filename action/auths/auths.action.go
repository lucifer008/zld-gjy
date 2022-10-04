package auths

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"zld-jy/models"
	"zld-jy/service/auths"
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
	loginUsers := auths.Instance.Auths(users)
	c.JSONP(http.StatusOK, loginUsers)
	log.Printf(">>>>>解析后的参数:", users)

}
func Authorize() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestURL := context.Request.RequestURI
		if strings.Contains(requestURL, "/login") {
			context.Next()
		}
		token := context.GetHeader("accesstoken")
		if len(token) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, "非法请求")
		}

	}
}
