package pubsub

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"

	"github.com/anant-sharma/go-blockchain/config"
	"github.com/anant-sharma/go-utils/mq"
)

// PubSub Struct
type PubSub struct {
	exchange string
	queue    string
	mq       mq.MQ
}

// Message Structure
type Message struct {
	Event string
	Data  interface{}
}

var _mq = mq.NewMQ()
var _pubsub PubSub

// NewPubSub Function
func NewPubSub(exchange string, queue string) {

	// Connect to MQ
	_mq.Connect(config.GetConfig().MQConnectionString)

	// Create Channel with MQ
	_mq.CreateChannel()

	// Create Exchange
	_mq.CreateExchange(exchange, "fanout", true)

	// Create Queue
	_mq.CreateQueue(queue, true, false, true, false, amqp.Table{})
	log.Printf("[*] Queue %s Created.", queue)

	// Bind Queue With Exchange
	_mq.BindQueueWithExchange(queue, exchange, "", amqp.Table{})
	log.Printf("[*] Queue %s Bound To Exchange %s.", queue, exchange)

	// Create Exchange Log Queue
	_mq.CreateQueue(exchange+"-logs", true, false, false, false, amqp.Table{
		"x-message-ttl": int(86400),
	})

	// Bind Exchange Log Queue
	_mq.BindQueueWithExchange(exchange+"-logs", exchange, "", amqp.Table{
		"message-ttl": int64(86400),
	})

	_pubsub = PubSub{
		exchange: exchange,
		queue:    queue,
	}
}

// Publish Method
func Publish(data Message) {
	_mq.Publish(_pubsub.exchange, _pubsub.queue, data)
}

// Subscribe Method
func Subscribe() <-chan Message {
	messages := _mq.EstablishWorker(_pubsub.queue)

	channel := make(chan Message)

	go func() {
		for d := range messages {
			if d.Body != nil {
				var j Message
				json.Unmarshal(d.Body, &j)

				go func() {
					channel <- j
				}()

				log.Printf(" [x] %s", d.Body)
			}
		}
	}()

	log.Printf("[*] Subscribed to queue %s", _pubsub.queue)
	return channel
}
