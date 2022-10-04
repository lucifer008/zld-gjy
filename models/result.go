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

//func OK(context *gin.Context) {
//	context.JSON(http.StatusOK, Result{
//		Code:    http.StatusOK,
//		Message: "success",
//	})
//}
