package tests

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
)

func Test_Token_Generator(t *testing.T) {
	var username = "zhangsan"
	var userId = 121
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"userId":   userId,
		"nbf":      time.Date(2022, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	//密钥
	var keyVal interface{}
	keyVal = []byte("123456")
	tokenString, err := token.SignedString(keyVal)
	fmt.Println(tokenString, err)
}
