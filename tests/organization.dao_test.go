package tests

import (
	"fmt"
	"testing"
	"zld-jy/da/model"
	"zld-jy/da/query"
)

func Test_Init_Org(t *testing.T) {
	qur := query.Use(DB)
	com, error := qur.Company.WithContext(ctx).First()
	if error != nil {
		fmt.Printf("%w", error)
		return
	}

	qur.Organization.WithContext(ctx).Create(&model.Organization{ID: 1, ComID: com.ID, OrgName: "zld", OrgCode: "zld-001", OrgLevel: 1})
}
