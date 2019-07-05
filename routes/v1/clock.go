package v1router

import (
	v1controller "github.com/anant-sharma/go-blockchain/controller/v1"
	"github.com/gin-gonic/gin"
)

// InitClockRouter - Function to initialize clock router
func InitClockRouter(router *gin.RouterGroup) {
	c1 := v1controller.NewController()

	router.GET("", c1.GetTimeStamp)
}
