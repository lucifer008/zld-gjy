package routers

import (
	"github.com/gin-gonic/gin"
	"zld-jy/action"
	"zld-jy/models"
)

func addCustomerRouters(rg *gin.RouterGroup) {
	rg.GET("/query", func(context *gin.Context) {
		var customer models.Customers
		context.ShouldBindQuery(&customer)
		total, data := action.CustomerActions.Query(customer)
		var result = models.ToResultPages(customer.Pages, total, data)
		models.OK(context, result)
	})
}
