package tests

import (
	"fmt"
	"log"
	"testing"
	"zld-jy/da/domain"
)

func init() {
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold:             time.Second,   // Slow SQL threshold
	//		LogLevel:                  logger.Silent, // Log level
	//		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
	//		Colorful:                  false,         // Disable color
	//	},
	//)
	log.Println("init....................")
}

//func TestMain(t *testing.M) {
//	log.Println("main....................")
//}

//重新映射实体
func TestAdvanceQuery(t *testing.T) {
	testing.Init()
	type CustomerUser struct {
		id       int64
		UserName string
	}
	//uy := dao.Use(DB)
	var customerUser = &CustomerUser{}
	DB.Model(&domain.SysUser{}).Find(customerUser)
	fmt.Println(customerUser)
}

func TestRawQuery(t *testing.T) {
	type Result struct {
		id       int
		empNo    string
		empName  string
		username string
	}
	var sql = "SELECT usr.Id as id,usr.User_Name as username FROM Sys_Users usr inner join Employee emp on usr.Emp_Id=emp.Id where usr.User_Name=?"
	var res = Result{}
	//todo 验证未通过
	DB.Raw(sql, "zhangsan").Scan(&res)
	fmt.Println(res)
}
