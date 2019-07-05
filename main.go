package main

import (
	"github.com/anant-sharma/go-blockchain/controller/v1/blockchain"
	router "github.com/anant-sharma/go-blockchain/routes"
)

func main() {

	blockchain.CreateNewBlockchain()

	router.InitRouter()
}
