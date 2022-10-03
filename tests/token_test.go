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
func Test_Token_Check(t *testing.T) {
	var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE2NjU0MDMyMDAsInVzZXJJZCI6MTIxLCJ1c2VybmFtZSI6InpoYW5nc2FuIn0.rPyzZssFatQr5_ZTnVVTSV3WfE8_-5o0EXuii0niHBI"
	//var hmacSampleSecret map[string]string
	//密钥
	var keyVal interface{}
	keyVal = []byte("123456")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		token.Valid = true
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return keyVal, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["username"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
}
