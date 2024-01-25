package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)}
	nonce := 0
	for {
		block.MyBlockHash = CalculateHash(*block, nonce)
		if isValid(block.MyBlockHash) {
			break
		}
		nonce++
	}
	return block
}

func CalculateHash(block Block, nonce int) []byte {
	data := []byte(fmt.Sprintf("%d%s%s%d", block.Timestamp, block.PreviousBlockHash, block.AllData, nonce))
	hash := sha256.Sum256(data)
	return hash[:]
}

func isValid(hash []byte) bool {
	return bytes.HasPrefix(hash, []byte{0, 0, 0})
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
