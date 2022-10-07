package support

import (
	"github.com/go-eden/routine"
	"zld-jy/models"
)

var currentContext = routine.NewLocalStorage()
var UserContextInstance UserContext

func init() {
	UserContextInstance = UserContext{}
}

type UserContext struct {
}

func (uc UserContext) SetCurrentUser(user models.CurrentUsers) {
	currentContext.Set(user)
}
func (uc UserContext) GetCurrentUser() models.CurrentUsers {
	return currentContext.Get().(models.CurrentUsers)
}
