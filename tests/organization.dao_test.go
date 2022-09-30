package tests

import (
	"fmt"
	"testing"
	"time"
	"zld-jy/da/model"
	"zld-jy/da/query"
)

func Test_Init_Org(t *testing.T) {
	qur := query.Use(DB)
	com, error := qur.Company.WithContext(ctx).First()
	if error != nil {
		fmt.Errorf("%w", error)
		return
	}
	org, _ := qur.Organization.WithContext(ctx).Count()
	newOrgId := org + 1
	err := qur.Organization.WithContext(ctx).Create(&model.Organization{ID: newOrgId,
		ComID: com.ID, OrgName: "zld", OrgCode: "zld-001", OrgLevel: 1, BeginDate: time.Now(), EndDate: time.Now().AddDate(0, 1, 0), InsertUser: 1, InsertDateTime: time.Now(), UpdateUser: 1, UpdateDateTime: time.Now(), Version: 0, Deleted: "0"})
	if err != nil {
		fmt.Errorf("错误>>>>>>%w", err)
		panic(err.Error())
		//return
	}
	fmt.Println(">>>>>>>>>>>>>>>create org sucess!")
}
