package tests_test

import (
	"testing"
	"time"
	"zld-jy/da/model"
	"zld-jy/da/query"
)

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
