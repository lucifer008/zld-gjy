package tests

import (
	"testing"
	"time"
	"zld-jy/da/model"
	"zld-jy/da/query"
)

//init employee
func Test_InitEmployee(t *testing.T) {
	dao := query.Use(DB)
	empId, _ := dao.Employee.WithContext(ctx).Count()
	empId = empId + 1
	org, _ := dao.Organization.WithContext(ctx).First()
	orgId := org.ID
	dao.Employee.WithContext(ctx).Create(&model.Employee{ID: empId, OrgID: orgId, EmpNo: "001", EmpName: "张三", EmpMobile: "18211329010", InsertUser: 1, InsertDateTime: time.Now(), UpdateDateTime: time.Now(), UpdateUser: 0, Version: 0, Deleted: "0"})

}
