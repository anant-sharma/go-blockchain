package v1router

import (
	"net/http"

	"github.com/anant-sharma/go-blockchain/controller/v1/blockchain"
	"github.com/gin-gonic/gin"
)

// InitBlockchainRouter - Function to initialize blockchain router
func InitBlockchainRouter(router *gin.RouterGroup) {
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, blockchain.B)
	})
}
