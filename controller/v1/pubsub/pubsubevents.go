package pubsub

type pubsubEvents struct {
	BlockMined         string
	TransactionCreated string
	ChainPublished     string
	ChainRequested     string
}

// PubSubEvents Constants
var PubSubEvents = pubsubEvents{
	BlockMined:         "BLOCK.MINED",
	TransactionCreated: "TRANSACTION.CREATED",
	ChainPublished:     "CHAIN.PUBLISHED",
	ChainRequested:     "CHAIN.REQUESTED",
}
