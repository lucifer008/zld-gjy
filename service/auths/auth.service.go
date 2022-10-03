package auths

import (
	"context"
	"zld-jy/da/base"
	"zld-jy/da/query"
	"zld-jy/models"
	"zld-jy/utils"
)

var Instance *AuthServiceImpl

func init() {
	Instance = &AuthServiceImpl{}
}

type AuthService interface {
	Auths(u models.Users) models.Auths
}
type AuthServiceImpl struct {
}

func (impl *AuthServiceImpl) Auths(u models.Users) models.Auths {
	qur := query.Use(base.DB)
	userDao := qur.SysUser
	employeeDao := qur.Employee
	var username = u.Username
	var password = utils.MD5(u.Password)
	userinfo, _ := userDao.WithContext(context.Background()).Where(userDao.UserName.Eq(username), userDao.UserPassword.Eq(password)).Take()
	if userinfo == nil {
		return models.Auths{Status: 1, Desc: "无此用户信息"}
	}
	employee, _ := employeeDao.WithContext(context.Background()).Where(employeeDao.ID.Eq(userinfo.EmpID)).Take()

	return models.Auths{Username: userinfo.UserName, EmplyeeName: employee.EmpName, AuthsToken: "2222222222222222", Status: 0, Desc: "成功!"}
}
