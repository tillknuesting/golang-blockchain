package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

// Transaction represents a simple transaction with a sender, receiver, and an amount.
type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
}

// Block represents a block in the blockchain.
type Block struct {
	Index        int
	Timestamp    int64
	Transactions []Transaction
	PrevHash     string
	Hash         string
}

// Blockchain is a series of blocks.
type Blockchain struct {
	Blocks []*Block
}

// NewBlock creates a new block with the given transactions.
func NewBlock(index int, transactions []Transaction, prevHash string) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     prevHash,
	}

	block.Hash = block.calculateHash()
	return block
}

// serializeTransactions serializes the transactions in a block ensuring that it works correctly even when the block is empty.
func (b *Block) serializeTransactions() string {
	serialized, err := json.Marshal(b.Transactions)
	if err != nil {
		return "[]"
	}
	return string(serialized)
}

// calculateHash generates a SHA-256 hash for the block.
func (b *Block) calculateHash() string {
	serializedTransactions := b.serializeTransactions()
	blockData := fmt.Sprintf("%d%d%s%s", b.Index, b.Timestamp, serializedTransactions, b.PrevHash)
	hash := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hash[:])
}

// AddBlock appends a new block with the given transactions to the blockchain.
func (bc *Blockchain) AddBlock(transactions []Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, transactions, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockchain creates a new blockchain with a genesis block.
func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock(0, []Transaction{}, "0")
	return &Blockchain{Blocks: []*Block{genesisBlock}}
}

func main() {
	// Initialize a new blockchain.
	blockchain := NewBlockchain()

	// Add a block with a sample transaction.
	blockchain.AddBlock([]Transaction{
		{
			Sender:   "Alice",
			Receiver: "Bob",
			Amount:   50,
		},
	})

	// Add a block with a sample transaction.
	blockchain.AddBlock([]Transaction{
		{
			Sender:   "Bob",
			Receiver: "Alice",
			Amount:   25,
		},
	})

	// Print the blocks in the blockchain.
	for _, block := range blockchain.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Prev. hash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("Transactions:")
		for _, tx := range block.Transactions {
			fmt.Printf("  Sender: %s, Receiver: %s, Amount: %.2f\n", tx.Sender, tx.Receiver, tx.Amount)
		}
		fmt.Println()
	}
}
