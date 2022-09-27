package routers

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// Run will start the server
func Run() {
	getRoutes()
	router.Run(":5000")
}
func getRoutes() {
	user := router.Group("/users")
	order := router.Group("/orders")
	customer := router.Group("customers")
	quotes := router.Group("/quotes")
	auths := router.Group("/auths")
	finance := router.Group("/finance")
	commons := router.Group("/commons")
	stat := router.Group("/stats")
	addUserRouters(user)
	addOrderRouters(order)
	addCustomerRouters(customer)
	addQuoteRouters(quotes)
	addAuthsRouters(auths)
	addFinanceRouters(finance)
	addCommonsRouters(commons)
	addStatsRouters(stat)
}
