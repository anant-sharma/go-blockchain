package blockchain

import (
	utils "github.com/anant-sharma/go-blockchain/common"
	"github.com/anant-sharma/go-blockchain/controller/v1/pubsub"
	"github.com/mitchellh/mapstructure"
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

	// Init PubSub
	pubsub.NewPubSub("bc.msg.exchange", utils.GenerateUUID())

	// Add Genesis Block
	B.NewBlock(11, utils.Sha256("previousBlockHash"), utils.Sha256("hash"))

	go func() {
		// Add Event Subscribers
		messages := pubsub.Subscribe()

		for msg := range messages {

			switch msg.Event {
			case pubsub.PubSubEvents.TransactionCreated:
				{
					var transaction Transaction
					mapstructure.Decode(msg.Data, &transaction)
					B.AddTransactionToPendingTransactions(transaction)
				}
			case pubsub.PubSubEvents.BlockMined:
				{
					var block Block
					mapstructure.Decode(msg.Data, &block)
					B.AddMinedBlockToChain(block)
				}
			}
		}
	}()

	go func() {
		B.RequestChain()
	}()

	return B
}
