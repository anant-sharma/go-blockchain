package v1router

import (
	"encoding/json"
	"net/http"

	"github.com/anant-sharma/go-blockchain/controller/v1/blockchain"
	"github.com/gin-gonic/gin"
)

type transactionInput struct {
	Data         string
	DataCategory string
	Sender       string
	Recipient    string
}

// InitTransactionRouter function
func InitTransactionRouter(router *gin.RouterGroup) {
	router.POST("", func(ctx *gin.Context) {

		form, _ := ctx.GetRawData()
		var tx transactionInput
		json.Unmarshal(form, &tx)

		transaction := blockchain.NewTransaction(tx.Data, tx.DataCategory, tx.Sender, tx.Recipient)
		ctx.JSON(http.StatusOK, transaction)
	})
}
