package tests_test

import (
	"context"
	"sync"
	"zld-jy/da/query"
)

var useOnce sync.Once
var ctx = context.Background()

func CRUDInit() {
	//绑定数据库
	query.Use(DB)

	//初始化对指针类型
	query.SetDefault(DB)
}
