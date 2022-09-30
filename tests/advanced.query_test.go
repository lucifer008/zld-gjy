package tests

import (
	"fmt"
	"testing"
	"zld-jy/da/model"
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
}

//重新映射实体
func TestAdvanceQuery(t *testing.T) {
	type CustomerUser struct {
		id       int64
		UserName string
	}
	//uy := query.Use(DB)
	var customerUser = &CustomerUser{}
	DB.Model(&model.SysUser{}).Find(customerUser)
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
	DB.Raw(sql, "zhangsan").Scan(&res).Commit()
	fmt.Println(res)
}
