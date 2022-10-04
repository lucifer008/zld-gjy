package action_users

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

// Login
//@Summary 获取用户信息
// @Schemes
// @Description 根据用户ID获取用户信息
// @Tags ops 认证
// @Param user body userId true "登录参数"
// @Accept json
// @Produce json
// @Success 200
// @Router /users/getUsers [get]
func (uh *UsersAction) GetUserInfo(userId string) models.UsersInfo {
	//log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>UserAction>>>>>GetUserInfo>>>>>>>>>>>")
	us, em := service_users.Instance.GetUsers(userId)
	return models.UsersInfo{UserId: us.ID, UserName: us.UserName, UserEmail: us.UserEmail, UserType: us.UserType, EmpNo: em.EmpNo, EmpName: em.EmpName}
}
