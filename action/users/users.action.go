package action_users

import (
	"github.com/gin-gonic/gin"
	"log"
	"zld-jy/service/users"
)

var Instance *UsersAction

func init() {
	Instance = &UsersAction{}
}

type UsersAction struct {
}

func (uh *UsersAction) GetUserInfo(c *gin.Context) {
	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>UserAction>>>>>GetUserInfo>>>>>>>>>>>")
	service_users.Instance.GetUsers()
}
