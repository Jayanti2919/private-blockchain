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

	filePath := "./xyz.txt"

	cid, err := FileUpload(filePath)

	if err != nil {
		fmt.Println(err)
		return
	}

	err2 := FileDownload(cid)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}
