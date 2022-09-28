package main

import (
	"fmt"
	"golang.org/x/net/context"
	"gorm.io/gen/examples/dal"
	"zld-jy/conf"
	"zld-jy/da/query"
	"zld-jy/routers"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()
	userDao := query.Use(dal.DB).SysUser
	user, err := userDao.WithContext(context.Background()).First()
	if err != nil {
		fmt.Printf("错误", err)
		return
	}
	fmt.Println(user)
}
func main() {
	routers.Run()
}
