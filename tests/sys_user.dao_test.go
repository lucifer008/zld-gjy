package tests

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"testing"
	"time"
	"zld-jy/da/model"
	"zld-jy/da/query"
)

func Test_InitUsers(t *testing.T) {
	dao := query.Use(DB)
	userId, _ := dao.SysUser.WithContext(ctx).Count()
	userId = userId + 1
	emp, _ := dao.Employee.WithContext(ctx).First()
	empId := emp.ID
	var password = MD5("wxllx@124343.com")
	error := dao.SysUser.WithContext(ctx).Create(&model.SysUser{ID: userId, EmpID: empId, UserName: "zhangsan", UserEmail: "zhangsan@163.com", UserPassword: password, UserStatus: 0, UserType: 1, InsertDateTime: time.Now(), InsertUser: 1, UpdateDateTime: time.Now(), UpdateUser: 1, Version: 1, Deleted: "0", CompayID: 1})
	if error != nil {
		panic(error)
	}
}
func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
func Test_LoginUser(t *testing.T) {
	sysUsers := query.Use(DB).SysUser
	var username = "zhangsan"
	password := MD5("wxllx@124343.com")
	count, err := sysUsers.WithContext(ctx).Where(sysUsers.UserName.Eq(username), sysUsers.UserPassword.Eq(password)).Count()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	log.Println(">>>>>>>>>>>>count=", count)
	//var count2 *int64
	//DB.Table("sys_users").Where("user_name=?", username).Where("user_password=?", password).Count(count2)

	//fmt.Println(count2)
}
