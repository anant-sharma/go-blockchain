package blockchain

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/anant-sharma/go-blockchain/controller/v1/pubsub"
	"github.com/mitchellh/mapstructure"

	utils "github.com/anant-sharma/go-blockchain/common"
	"github.com/anant-sharma/go-blockchain/config"
	"github.com/anant-sharma/go-utils/mq"
	"github.com/streadway/amqp"
)

type chainRequest struct {
	ReplyQueue string
}

var _mq = mq.NewMQ()

// RequestChain Method
func (b *Blockchain) RequestChain() {

	queue := utils.GenerateUUID()

	// Connect to MQ
	_mq.Connect(config.GetConfig().MQConnectionString)

	// Create Channel with MQ
	_mq.CreateChannel()

	// Create Queue
	_mq.CreateQueue(queue, true, true, false, false, amqp.Table{})
	log.Printf("[*] Queue %s Created.", queue)

	// Publish Chain Request
	pubsub.Publish(pubsub.Message{
		Data: chainRequest{
			ReplyQueue: queue,
		},
		Event: pubsub.PubSubEvents.ChainRequested,
	})

	// Add Listeners
	messages := _mq.EstablishWorker(queue)

	go func() {
		for d := range messages {
			if d.Body != nil {
				var j pubsub.Message
				json.Unmarshal(d.Body, &j)

				switch j.Event {
				case pubsub.PubSubEvents.ChainPublished:
					{
						var blockchain Blockchain
						mapstructure.Decode(j.Data, &blockchain)
						b.SynchroniseChain(blockchain)
						break
					}
				}

				log.Printf(" [x] %s", d.Body)
			}
		}
	}()

	channel := make(chan bool)

	<-channel
}

// SynchroniseChain Method
func (b *Blockchain) SynchroniseChain(chain Blockchain) {
	Chain := chain.Chain
	PendingTransactions := chain.PendingTransactions

	// Return if incoming chain is smaller than current chain
	if len(Chain) <= len(b.Chain) {
		if len(b.PendingTransactions) < len(PendingTransactions) && b.areTransactionsValid(PendingTransactions) {
			b.PendingTransactions = PendingTransactions
		}
		return
	}

	// Check if incoming chain is valid
	if !isChainValid(Chain) {
		return
	}

	isValid := true
	for i := 0; i < len(b.Chain) && isValid; i++ {
		// Check if blocks are equal
		if !reflect.DeepEqual(b.Chain[i], Chain[i]) {
			isValid = false
			break
		}
	}

	if isValid {
		b.Chain = Chain
		b.PendingTransactions = PendingTransactions
	}
}

func isChainValid(chain []Block) bool {
	isValid := true

	for i := 1; i < len(chain) && isValid; i++ {
		block := chain[i]
		previousBlock := chain[i-1]

		// Hash Matching
		if previousBlock.Hash != block.Hash {
			isValid = false
			break
		}

		// Check Current Block Hash
		if (block.Hash != HashBlock(previousBlock.Hash, BlockData{
			Index:        i,
			Transactions: block.Transactions,
		}, block.Nonce)) {
			isValid = false
			break
		}
	}

	return isValid
}

func (b *Blockchain) areTransactionsValid(transactions []Transaction) bool {
	isValid := true

	for i := 0; i < len(b.PendingTransactions) && isValid; i++ {
		if !reflect.DeepEqual(b.PendingTransactions[i], transactions[i]) {
			isValid = false
			break
		}
	}

	return isValid
}
