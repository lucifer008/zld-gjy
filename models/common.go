package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code    int         `json:"code" `
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func OK(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, Result{
		Code:    http.StatusOK,
		Data:    data,
		Message: "success",
	})
}

type Pages struct {
	PageSize     int
	CurrentIndex int
	Total        int64
	Data         interface{}
}
type Base struct {
	PageParams Pages
}

func ToPages(pageSize int, currentIndex int, total int64, data interface{}) Pages {
	return Pages{pageSize, currentIndex, total, data}
}

//func OK(context *gin.Context) {
//	context.JSON(http.StatusOK, Result{
//		Code:    http.StatusOK,
//		Message: "success",
//	})
//}
