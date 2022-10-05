package tests

import (
	"strconv"
	"strings"
	"testing"
	"zld-jy/da/model"
	"zld-jy/da/query"
)

func Test_InitCustomers(t *testing.T) {
	qur := query.Use(DB)
	count, _ := qur.Customer.WithContext(ctx).Count()
	var rows = 10
	for index := 1; index < rows; index++ {
		var id = count + 1
		qur.Customer.WithContext(ctx).Create(&model.Customer{
			ID:              id,
			CustomerName:    "张三公司" + strconv.FormatInt(id, 10),
			CustomerAddress: "西安市未央区凤城十二路广场" + strconv.FormatInt(id, 10),
			CompayID:        1,
			ContactMan:      strings.Join([]string{"张三", strconv.FormatInt(id, 10)}, "-"),
			ContactTel:      "1871132901" + strconv.Itoa(index),
		})
	}
}
