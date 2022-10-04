package auths

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"strings"
	"zld-jy/config"
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
		success := checkAuths(token)
		if success {
			context.JSON(http.StatusUnauthorized, "Token已过期或者无效Token")
		}
	}
}
func checkAuths(tokenString string) bool {
	var keyVal interface{}
	keyVal = []byte(config.Config.Token.Screct)
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		token.Valid = true
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return keyVal, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(">>>>>>Token认证信息>>>>>>>", claims["userId"], claims["username"], claims["nbf"])
		return true
	}
	return false
}
