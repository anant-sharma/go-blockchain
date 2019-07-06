package v1router

import (
	"net/http"

	"github.com/anant-sharma/go-blockchain/controller/v1/blockchain"
	"github.com/gin-gonic/gin"
)

// InitTransactionRouter function
func InitTransactionRouter(router *gin.RouterGroup) {
	router.POST("", func(ctx *gin.Context) {
		transaction := blockchain.NewTransaction(100, "a", "b")
		ctx.JSON(http.StatusOK, transaction)
	})
}
