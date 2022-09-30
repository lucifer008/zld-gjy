package tests_test

import (
	"fmt"
	"testing"
	"time"
	"zld-jy/da/model"
	"zld-jy/da/query"
)

//测试create
func Test_newCompany(t *testing.T) {
	useOnce.Do(CRUDInit)

	u := query.Company

	id, _ := u.WithContext(ctx).Count()
	err := u.WithContext(ctx).Create(&model.Company{
		ID:          id + 1,
		CompanyName: "测试", RegisterDateTime: time.Now(),
		InsertUser:     1,
		InsertDateTime: time.Now(),
		UpdateUser:     1,
		UpdateDateTime: time.Now(),
		Version:        1,
		Deleted:        "0",
	})
	if err != nil {
		t.Errorf("create model fail: %s", err)
	}
}

//测试事务
func Test_Trans(t *testing.T) {
	qur := query.Use(DB)
	err := qur.Transaction(func(tx *query.Query) error {
		id, _ := tx.Company.WithContext(ctx).Count()
		ids := id + 1
		fmt.Println("插入Id", ids)
		tx.Company.WithContext(ctx).Create(&model.Company{
			ID:          id + 1,
			CompanyName: "测试", RegisterDateTime: time.Now(),
			InsertUser:     1,
			InsertDateTime: time.Now(),
			UpdateUser:     1,
			UpdateDateTime: time.Now(),
			Version:        1,
			Deleted:        "0",
		})
		if ids > 1 {
			panic("手动异常，测试回滚!")
		}
		tx.Company.WithContext(ctx).Delete(&model.Company{ID: id + 11})

		return nil
	})
	if err != nil {
		fmt.Errorf("错误:%w", err)
	}
	fmt.Println(">>>>>>>>>>>>>success!>>>>>>>>>>>>>>>>>>>>>>>>>")
}

//测试list
func Test_Company_Query(t *testing.T) {
	qur := query.Use(DB)
	//var comList []*model.Company
	comList, err := qur.Company.WithContext(ctx).Find()
	if err != nil {
		fmt.Errorf("错误>>>>%w", err)
	}
	for com, v := range comList {
		fmt.Println(com, v)
	}
	//fmt.Println(comList)
}
