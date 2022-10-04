package models

type LoginUsers struct {
	Username string `json:"username" example:"zhangsan"`         //用户名
	Password string `json:"password" example:"wxllx@124343.com"` //密码
}
type UsersInfo struct {
	UserId    int64  //用户Id
	UserName  string //用户名称
	UserEmail string //用户邮箱
	UserType  int32  //用户类型
	EmpNo     string //雇员编号
	EmpName   string //雇员名称
}
type RegisterUser struct {
	UserName  string //用户名称
	UserEmail string //用户邮箱
	Tel       string //手机号
	Password  string //密码
}
