package tests

import (
	"context"
	"sync"
	"zld-jy/da/dao"
)

var useOnce sync.Once
var ctx = context.Background()

func CRUDInit() {
	//绑定数据库
	dao.Use(DB)

	//初始化对指针类型
	dao.SetDefault(DB)
}
