package blockchain

import (
	utils "github.com/anant-sharma/go-blockchain/common"
)

// B - The Blockchain
var B Blockchain

// Blockchain structure
type Blockchain struct {
	Chain               []Block
	PendingTransactions []Transaction
}

// CreateNewBlockchain - Function to create new chain
func CreateNewBlockchain() Blockchain {
	B = Blockchain{
		Chain:               make([]Block, 0),
		PendingTransactions: make([]Transaction, 0),
	}

	// Add Genesis Block
	B.NewBlock(11, utils.Sha256("previousBlockHash"), utils.Sha256("hash"))

	return B
}
