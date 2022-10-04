package models

type LoginUsers struct {
	Username string
	Password string
}
type UsersInfo struct {
	UserId    int64  //用户Id
	UserName  string //用户名称
	UserEmail string //用户游戏
	UserType  int32  //用户类型
	EmpNo     string //雇员编号
	EmpName   string //雇员名称
}
