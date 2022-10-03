package main

import (
	"fmt"
	"zld-jy/conf"
	"zld-jy/da/base"
	"zld-jy/routers"
)

func init() {
	base.DB = base.ConnectDB(conf.MySQLDSN).Debug()
	//userDao := query.Use(base.DB).SysUser
	//user, err := userDao.WithContext(context.Background()).First()
	//if err != nil {
	//	fmt.Printf("错误", err)
	//	return
	//}
	//fmt.Printf(">>>>>>>>>查询到信息:%+v\n", user)
	fmt.Println(">>>>>>>>>>数据库初始化成功!>>>>>>>>>>>>")
}
func main() {
	routers.Run()
}
