package main

func (blockchain *Blockchain) AddBlock(data string) {
	PrevBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newblock := NewBlock(data, PrevBlock.MyBlockHash)
	blockchain.Blocks = append(blockchain.Blocks, newblock)
}

// adds genesis block as the first block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
