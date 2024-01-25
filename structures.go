package main

type Block struct {
	Timestamp         int64
	PreviousBlockHash []byte
	MyBlockHash       []byte
	AllData           []byte
	Nonce             int
}

type Blockchain struct {
	Blocks []*Block
}
