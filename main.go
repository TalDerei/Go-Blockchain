package main

import (
	"github.com/TalDerei/Go-Blockchain/blockchain"
	"fmt"
	"strconv"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First block after genisis block!")
	chain.AddBlock("Second block after genisis block!")
	chain.AddBlock("Third block after genisis block!")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block : %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}