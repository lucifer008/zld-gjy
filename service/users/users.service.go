package service_users

import (
	"context"
	"strconv"
	"zld-jy/da/base"
	"zld-jy/da/dao"
	"zld-jy/da/domain"
)

type UsersService interface {
	GetUsers()
}

var Instance *UserServiceImpl

func init() {
	Instance = &UserServiceImpl{}
}

type UserServiceImpl struct {
}

func (u *UserServiceImpl) GetUsers(userId string) (us *domain.SysUser, em *domain.Employee) {
	if userId == "" {
		panic("用户Id不能为空!")
	}
	qur := dao.Use(base.DB)
	sysUserQuery := qur.SysUser
	employeeQuery := qur.Employee
	id, _ := strconv.ParseInt(userId, 10, 64)
	sysUser, _ := sysUserQuery.WithContext(context.Background()).Where(sysUserQuery.ID.Eq(id)).Take()
	//if err != nil {
	//	panic(err)
	//}
	employee, _ := employeeQuery.WithContext(context.Background()).Where(employeeQuery.ID.Eq(sysUser.EmpID)).Take()
	return sysUser, employee
}
