package tests_test

import (
	"context"
	"sync"
	"zld-jy/da/query"
)

var useOnce sync.Once
var ctx = context.Background()

func CRUDInit() {
	query.Use(DB)
	query.SetDefault(DB)
}
