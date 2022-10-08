package action

import (
	"context"
	"errors"
	"log"
	"zld-jy/da/base"
	"zld-jy/da/dao"
	"zld-jy/da/domain"
	"zld-jy/models"
	"zld-jy/service/users"
	"zld-jy/support"
	"zld-jy/utils"
)

var UserActionInstance *UsersAction

func init() {
	UserActionInstance = &UsersAction{}
}

//UsersAction 用户相关
type UsersAction struct {
	BaseAction
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
	log.Println(">>>>>>>>>>>>>>>>当前用户信息>>>>>>>>>>>>>>>", support.UserContextInstance.GetCurrentUser())
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
func (ua *UsersAction) Register(users models.RegisterUser) error {
	qur := dao.Use(base.DB)
	if users.UserName == "" {
		return errors.New("用户名不能为空")
	}
	if users.Password == "" {
		return errors.New("密码不能为空")
	}
	if users.UserEmail == "" {
		return errors.New("邮箱不能为空")
	}
	if users.Tel == "" {
		return errors.New("手机号不能为空")
	}
	count, _ := qur.SysUser.WithContext(context.Background()).Or(qur.SysUser.UserName.Eq(users.UserName)).Or(qur.SysUser.UserEmail.Eq(users.UserEmail)).Count()
	if count == 0 {
		count, _ = qur.Employee.WithContext(context.Background()).Where(qur.Employee.EmpMobile.Eq(users.Tel)).Count()
		if count > 0 {
			return errors.New("手机号已存在")
		}
	}
	if count > 0 {
		return errors.New("用户名或者邮箱已存在")
	}

	qur.Transaction(func(tx *dao.Query) error {
		var empId, _ = qur.Employee.WithContext(context.Background()).Count()
		var employee domain.Employee
		employee = domain.Employee{
			ID:             empId + 1,
			OrgID:          orgId,
			EmpNo:          "001",
			EmpMobile:      users.Tel,
			EmpName:        users.UserName,
			InsertDateTime: insertDatetime,
			InsertUser:     insertUser,
			UpdateUser:     updateUser,
			UpdateDateTime: updateDatetime,
			Deleted:        deleted,
			Version:        version,
		}
		err := qur.Employee.WithContext(context.Background()).Create(&employee)
		if err != nil {
			panic(err)
		}
		var userId, _ = qur.SysUser.WithContext(context.Background()).Count()
		var sysUsers domain.SysUser
		sysUsers = domain.SysUser{ID: userId + 1,
			EmpID:          empId,
			UserName:       users.UserName,
			UserEmail:      users.UserEmail,
			UserPassword:   utils.MD5(users.Password),
			UserStatus:     0,
			UserType:       1,
			InsertUser:     insertUser,
			InsertDateTime: insertDatetime,
			UpdateUser:     updateUser,
			UpdateDateTime: updateDatetime,
			Deleted:        deleted,
			Version:        version,
			CompayID:       companyId,
		}
		err = qur.SysUser.WithContext(context.Background()).Create(&sysUsers)
		if err != nil {
			panic(err)
		}
		return nil
	})
	return nil
}
