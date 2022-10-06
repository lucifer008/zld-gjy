package support

import (
	"github.com/go-eden/routine"
	"zld-jy/models"
)

var currentContenxt = routine.NewLocalStorage()
var UserContextInstance UserContext

func init() {
	UserContextInstance = UserContext{}
}

type UserContext struct {
}

func (uc UserContext) SetCurrentUser(user models.CurrentUsers) {
	currentContenxt.Set(user)
}
func (uc UserContext) GetCurrentUser() models.CurrentUsers {
	return currentContenxt.Get().(models.CurrentUsers)
}
