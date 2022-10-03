package tests

import (
	"fmt"
	"testing"
	"zld-jy/middleware"
)

func Test_Redis(t *testing.T) {
	middleware.Set("test", "testsss")
}
func Test_Redis_Get(t *testing.T) {
	fmt.Println(middleware.Get("test"))
}
