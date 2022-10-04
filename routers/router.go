package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"zld-jy/config"
	"zld-jy/docs"
)

var router = gin.Default()

// Run will start the server
func Run() {
	log.Println(">>>>>>>start service.........")
	docs.SwaggerInfo.BasePath = "/api"
	getRoutes()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	log.Println(">>>>>>>router 加载成功.........")

	router.Run(config.Config.Server.Port)

}
func getRoutes() {
	user := router.Group("/api/users")
	order := router.Group("/api/orders")
	customer := router.Group("/api/customers")
	quotes := router.Group("/api/quotes")
	auths := router.Group("/api/auths")
	finance := router.Group("/api/finance")
	commons := router.Group("/api/commons")
	stat := router.Group("/api/stats")
	addUserRouters(user)
	addOrderRouters(order)
	addCustomerRouters(customer)
	addQuoteRouters(quotes)
	addAuthsRouters(auths)
	addFinanceRouters(finance)
	addCommonsRouters(commons)
	addStatsRouters(stat)
}
