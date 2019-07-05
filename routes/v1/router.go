package v1router

import (
	"github.com/gin-gonic/gin"
)

// InitRouter Function
func InitRouter(v1 *gin.RouterGroup) {

	InitClockRouter(v1.Group("/clock"))
	InitBlockchainRouter(v1.Group("/blockchain"))
	InitMineRouter(v1.Group("/mine"))
	InitTransactionRouter(v1.Group("transactions"))

}
