package v1router

import (
	"net/http"

	"github.com/anant-sharma/go-blockchain/controller/v1/blockchain"
	"github.com/gin-gonic/gin"
)

// InitMineRouter function
func InitMineRouter(router *gin.RouterGroup) {
	router.GET("", func(ctx *gin.Context) {
		NewBlock := blockchain.B.MineBlock()
		ctx.JSON(http.StatusOK, NewBlock)
	})
}
