package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []string
}

// Function to create NewBlock
// Function Name: NewBlock
// Inputs: nonce, previousHash
// Data Type of return Values: *Block (pointer to a Block structure)
func NewBlock(nonce int, previousHash [32]byte) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

// Method to print a Block
// Method Name: Print
// Example: block.Print() block being a pointer to a block structure
func (b *Block) Print() {
	fmt.Printf("timestamp			%d\n", b.timestamp)
	fmt.Printf("nonce			%d\n", b.nonce)
	fmt.Printf("previous_hash			%s\n", b.previousHash)
	fmt.Printf("transactions			%s\n", b.transactions)
}

// Method to get the hash of a block
// Method Name: Hash
// Example: block.Hash() block being a pointer to a block structure
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

// Method to transform a block to JSON format
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        int      `json:"nonce"`
		PreviousHash [32]byte `json:"previous_hash"`
		Transactions []string `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

// Function to create new blockchain
// Function Name: NewBlockchain
// Inputs:
// Data Type of return Values: *Blockchain (pointer to a Blockchain structure)
func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

// Method to create a block in a blockchain
// Method Name: CreateBlock
// Example: blockchain.CreateBlock() blockchain being a pointer to a blockchain structure
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

// Method to get the last block
// Method Name: LastBlock
// Example: blockchain.LastBlock() blockchain being a pointer to a blockchain structure
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

// Method to gprint a blockchain
// Method Name: Print
// Example: blockchain.Print() block being a pointer to a blockchain structure
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i,
			strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func main() {
	blockChain := NewBlockchain()
	blockChain.Print()
	fmt.Printf("\n---------------------------------------------------\n")

	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, previousHash)
	blockChain.Print()
	fmt.Printf("\n---------------------------------------------------\n")

	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, previousHash)
	blockChain.Print()
	fmt.Printf("\n---------------------------------------------------\n")
}
