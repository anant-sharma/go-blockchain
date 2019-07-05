package blockchain

import (
	utils "github.com/anant-sharma/go-blockchain/common"
)

// Transaction structure
type Transaction struct {
	Amount        float64
	Recipient     string
	Sender        string
	TransactionID string
}

// NewTransaction to create new transaction
func NewTransaction(amount float64, sender string, recipient string) Transaction {

	// this.pubsub.publish({
	// 	Data: transaction,
	// 	Event: PUBSUB_EVENTS.TRANSACTION.CREATED,
	// });

	return Transaction{
		Amount:        amount,
		Recipient:     recipient,
		Sender:        sender,
		TransactionID: utils.GenerateUUID(),
	}
}

// AddTransactionToPendingTransactions to add transaction to chain
func (b *Blockchain) AddTransactionToPendingTransactions(transaction Transaction) int {
	b.PendingTransactions = append(b.PendingTransactions, transaction)
	return b.GetLastBlock().Index + 1
}
