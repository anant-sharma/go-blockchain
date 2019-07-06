package blockchain

import (
	"encoding/json"
	"math"
	"strings"
	"time"

	utils "github.com/anant-sharma/go-blockchain/common"
	"github.com/anant-sharma/go-blockchain/controller/v1/pubsub"
)

// Block structure
type Block struct {
	Hash              string
	Index             int
	Nonce             int
	PreviousBlockHash string
	Timestamp         int64
	Transactions      []Transaction
}

// BlockData structure
type BlockData struct {
	Index        int
	Transactions []Transaction
}

// NewBlock - To create new block
func (b *Blockchain) NewBlock(nonce int, previousBlockHash string, hash string) Block {

	var timestamp int64
	if len(b.Chain) == 0 {
		timestamp = 0
	} else {
		timestamp = int64(math.Ceil(float64(time.Now().UnixNano() / 1000000)))
	}

	block := Block{
		Hash:              hash,
		Index:             len(b.Chain) + 1,
		Nonce:             nonce,
		PreviousBlockHash: previousBlockHash,
		Timestamp:         timestamp,
		Transactions:      make([]Transaction, 0),
	}

	b.PendingTransactions = make([]Transaction, 0)
	b.Chain = append(b.Chain, block)

	return block
}

// GetLastBlock of blockchain
func (b *Blockchain) GetLastBlock() Block {
	return b.Chain[len(b.Chain)-1]
}

// MineBlock of blockchain
func (b *Blockchain) MineBlock() Block {
	lastBlock := b.GetLastBlock()
	previousBlockHash := lastBlock.PreviousBlockHash

	currentBlockData := BlockData{
		Index:        lastBlock.Index + 1,
		Transactions: b.PendingTransactions,
	}

	nonce := ProofOfWork(previousBlockHash, currentBlockData)

	blockHash := HashBlock(previousBlockHash, currentBlockData, nonce)

	newBlock := b.NewBlock(nonce, lastBlock.Hash, blockHash)

	pubsub.Publish(pubsub.Message{
		Event: pubsub.PubSubEvents.BlockMined,
		Data:  newBlock,
	})

	return newBlock
}

// AddMinedBlockToChain add block
func (b *Blockchain) AddMinedBlockToChain(newBlock Block) {
	lastBlock := b.GetLastBlock()

	// Check if hashes matches
	isHashCorrect := lastBlock.Hash == newBlock.PreviousBlockHash

	isIndexCorrect := lastBlock.Index+1 == newBlock.Index

	if isHashCorrect && isIndexCorrect {
		b.Chain = append(b.Chain, newBlock)
		b.PendingTransactions = make([]Transaction, 0)
	}
}

// HashBlock - sha256 hash of block
// Ref: https://github.com/openblockchains/awesome-sha256/blob/master/hash.go
func HashBlock(previousBlockHash string, currentBlockData BlockData, nonce int) string {
	blockDataString, err := json.Marshal(currentBlockData)
	if err != nil {
		panic(err)
	}

	data := previousBlockHash + string(nonce) + string(blockDataString)

	return utils.Sha256(data)
}

// ProofOfWork Algorithm
func ProofOfWork(previousBlockHash string, currentBlockData BlockData) int {
	var nonce int
	var hash string

	hash = HashBlock(previousBlockHash, currentBlockData, nonce)

	for {
		if strings.HasPrefix(hash, "0000") {
			break
		}
		nonce = nonce + 1
		hash = HashBlock(previousBlockHash, currentBlockData, nonce)
	}

	return nonce
}
