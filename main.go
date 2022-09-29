package main

import (
	"fmt"
	"golang.org/x/net/context"
	"zld-jy/conf"
	"zld-jy/da/base"
	"zld-jy/da/query"
	"zld-jy/routers"
)

func init() {
	base.DB = base.ConnectDB(conf.MySQLDSN).Debug()
	userDao := query.Use(base.DB).SysUser
	user, err := userDao.WithContext(context.Background()).First()
	if err != nil {
		fmt.Printf("错误", err)
		return
	}
	fmt.Printf(">>>>>>>>>查询到信息:%+v\n", user)
}
func main() {
	routers.Run()
}
