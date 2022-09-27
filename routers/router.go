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
	addUserRouters(user)
	addOrderRouters(order)
	addCustomerRouters(customer)
	addQuoteRouters(quotes)
}
