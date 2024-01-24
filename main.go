package main

import (
	"fmt"
)

func main() {
	fmt.Println("Creating the first block")

	newblockchain := NewBlockchain()
	newblockchain.AddBlock("First Block")
	newblockchain.AddBlock("Second Block")

	for _, block := range newblockchain.Blocks {
		fmt.Printf("Hash of the block: %x\n", block.MyBlockHash)
		fmt.Printf("Hash of the previous block: %x\n", block.PreviousBlockHash)
		stringValue := string(block.AllData)
		fmt.Println("Data of the block: ", stringValue)
	}
}
