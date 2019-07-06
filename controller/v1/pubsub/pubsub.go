package pubsub

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/streadway/amqp"

	"github.com/anant-sharma/go-blockchain/common/mq"
	"github.com/anant-sharma/go-blockchain/config"
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

	buf := &bytes.Buffer{}
	err := binary.Read(buf, binary.BigEndian, &data)
	if err != nil {
		panic(err)
	}

	_mq.Publish(_pubsub.exchange, _pubsub.queue, []byte(buf.Bytes()))
}
