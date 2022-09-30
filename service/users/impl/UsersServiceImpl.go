package service_users

import "log"

var Instance *UserServiceImpl

func init() {
	Instance = &UserServiceImpl{}
}

type UserServiceImpl struct {
}

func (u *UserServiceImpl) GetUsers() {
	log.Println(">>>>>>>>>>>>>>>>User Serive GetUsers>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
}
