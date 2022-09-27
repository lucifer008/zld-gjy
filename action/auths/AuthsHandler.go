package auths

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"zld-gly/models"
)

type AuthsHandler struct {
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func (ah AuthsHandler) Login(c *gin.Context) {
	log.Printf("请求成功")
	var users models.Users
	c.IndentedJSON(http.StatusOK, users)
	log.Printf(">>>>>解析后的参数:", users)
}
