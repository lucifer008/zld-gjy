package action

import (
	"zld-jy/models"
	"zld-jy/service/users"
)

var Instance *UsersAction

func init() {
	Instance = &UsersAction{}
}

type UsersAction struct {
}

// GetUserInfo
//@Summary 获取用户信息
// @Tags 用户
// @Schemes
// @Description 根据用户ID获取用户信息
// @Param userId query  string true "登录参数" default(1)
// @Accept json
// @Produce json
// @Success 200 {object} models.UsersInfo
// @Router /users/getUsers [get]
func (uh *UsersAction) GetUserInfo(userId string) models.UsersInfo {
	//log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>UserAction>>>>>GetUserInfo>>>>>>>>>>>")
	us, em := service_users.Instance.GetUsers(userId)
	return models.UsersInfo{UserId: us.ID, UserName: us.UserName, UserEmail: us.UserEmail, UserType: us.UserType, EmpNo: em.EmpNo, EmpName: em.EmpName}
}

// Register
//@Summary 注册用户
// @Tags 用户
// @Schemes
// @Description 注册用户
// @Param user body  models.RegisterUser true "注册用户信息"
// @Accept json
// @Produce json
// @Success 200 {object} models.Result
// @Router /users/register [post]
func Register(ru models.RegisterUser) {

}
