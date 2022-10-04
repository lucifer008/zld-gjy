package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Status int
	Desc   string
	Data   interface{}
}

func OK(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, Result{
		Status: http.StatusOK,
		Data:   data,
		Desc:   "success",
	})
}
