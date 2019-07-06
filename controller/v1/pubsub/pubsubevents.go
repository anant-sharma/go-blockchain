package pubsub

type pubsubEvents struct {
	BlockMined         string
	TransactionCreated string
}

// PubSubEvents Constants
var PubSubEvents = pubsubEvents{
	BlockMined:         "BLOCK.MINED",
	TransactionCreated: "TRANSACTION.CREATED",
}
