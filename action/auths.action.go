package action

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"zld-jy/config"
	"zld-jy/models"
	"zld-jy/service/auths"
	"zld-jy/support"
)

//var UserActionInstance *AuthsAction
//
//func init() {
//	UserActionInstance = &AuthsAction{}
//}
//
//type AuthsAction struct {
//}

var nonAuthUrl = [2]string{"swagger", "login"}

// Login
//@Summary 登录接口
// @Schemes
// @Description 登录
// @Tags 认证
// @Param user body models.LoginUsers true "登录参数"
// @Accept json
// @Produce json
// @Success 200
// @Router /auths/login [post]
func Login(c *gin.Context) {
	var users models.LoginUsers
	if err := c.BindJSON(&users); err != nil {
		c.JSONP(http.StatusOK, gin.H{"status": "错误:" + err.Error()})
		return
	}
	loginUsers := auths.Instance.Auths(users)
	models.OK(c, loginUsers)

}
func Authorize() gin.HandlerFunc {
	return func(context *gin.Context) {
		if !config.Config.Token.Enable {
			context.Next()
			return
		}

		requestURL := context.Request.RequestURI
		for _, v := range nonAuthUrl {
			if strings.Contains(requestURL, v) {
				context.Next()
				return
			}
		}
		token := context.GetHeader("accesstoken")
		if len(token) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, models.Result{Code: http.StatusUnauthorized, Message: "非法请求"})
			return
		}
		success := checkAuths(token)
		if !success {
			context.Abort()
			context.JSON(http.StatusUnauthorized, models.Result{Code: http.StatusUnauthorized, Message: "Token已过期或者无效Token"})
			return
		}
		context.Next()
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
		currentUser := models.CurrentUsers{UserId: int64(claims["userId"].(float64)), UserName: claims["username"].(string)}
		support.UserContextInstance.SetCurrentUser(currentUser)
		return true
	}
	return false
}
func GlobalExceptonHandler(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf(">>>>>>>>>>>>>>>错误>>>>>>>>>>>>>>>: %v\n", r)
			debug.PrintStack()
			c.JSON(http.StatusOK, models.Result{
				Code:    1001,
				Message: "内部错误:" + errorToString(r),
			})
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
