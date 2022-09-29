package tests_test

import (
	"context"
	"gorm.io/gen/examples/dal/query"
	"sync"
)

var useOnce sync.Once
var ctx = context.Background()

func CRUDInit() {
	query.Use(DB)
	query.SetDefault(DB)
}
