package blockchain

import (
	utils "github.com/anant-sharma/go-blockchain/common"
	"github.com/anant-sharma/go-blockchain/controller/v1/pubsub"
)

// Transaction structure
type Transaction struct {
	Checksum      string
	Data          string
	DataCategory  string
	Recipient     string
	Sender        string
	TransactionID string
}

// NewTransaction to create new transaction
func NewTransaction(data string, dataCategory string, sender string, recipient string) Transaction {

	transaction := Transaction{
		Checksum:      utils.Sha256(data),
		Data:          data,
		DataCategory:  dataCategory,
		Recipient:     recipient,
		Sender:        sender,
		TransactionID: utils.GenerateUUID(),
	}

	go func() {
		pubsub.Publish(pubsub.Message{
			Event: pubsub.PubSubEvents.TransactionCreated,
			Data:  transaction,
		})
	}()

	return transaction
}

// AddTransactionToPendingTransactions to add transaction to chain
func (b *Blockchain) AddTransactionToPendingTransactions(transaction Transaction) int {
	b.PendingTransactions = append(b.PendingTransactions, transaction)
	return b.GetLastBlock().Index + 1
}
