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
	PageSize     int `form:"pageSize" json:"pageSize" example:"10"`
	CurrentIndex int `form:"currentIndex" json:"currentIndex" example:"1"`
}
type ResultPages struct {
	Pages
	Total int64
	Data  interface{}
}
type Base struct {
	Pages
}

func ToResultPages(page Pages, total int64, data interface{}) ResultPages {
	return ResultPages{
		Pages: page,
		Total: total, Data: data,
	}
}

//func OK(context *gin.Context) {
//	context.JSON(http.StatusOK, Result{
//		Code:    http.StatusOK,
//		Message: "success",
//	})
//}
