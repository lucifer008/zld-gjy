package service_users

import "log"

type UsersService interface {
	GetUsers()
}

var Instance *UserServiceImpl

func init() {
	Instance = &UserServiceImpl{}
}

type UserServiceImpl struct {
}

func (u *UserServiceImpl) GetUsers() {
	log.Println(">>>>>>>>>>>>>>>>User Serive GetUsers>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
}
