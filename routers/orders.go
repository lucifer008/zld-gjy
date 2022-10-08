package routers

import (
	"github.com/gin-gonic/gin"
	"zld-jy/action"
	"zld-jy/models"
)

func addOrderRouters(rg *gin.RouterGroup) {
	rg.GET("/query", func(context *gin.Context) {
		var orders models.Orders
		context.ShouldBindQuery(&orders)
		total, data, err := action.OrderActionInstance.Query(orders)
		if err != nil {
			panic(err)
		}
		var result = models.ToResultPages(orders.Pages, total, data)
		models.OK(context, result)
	})
}
