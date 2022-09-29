package tests_test

import (
	"testing"
	"zld-jy/da/model"
	"zld-jy/da/query"
)

func Test_newCompany(t *testing.T) {
	useOnce.Do(CRUDInit)

	u := query.Company

	err := u.WithContext(ctx).Create(&model.Company{ID: 1, CompanyName: "测试"})
	if err != nil {
		t.Errorf("create model fail: %s", err)
	}
}
