
package blockchain

type BlockChain struct {
	// array of pointers to Block structs
	Blocks []*Block    
}

type Block struct {
	// byte (unsigned 8-bit integers) slices  
	Hash 		[]byte 
	Data 		[]byte
	PrevHash	[]byte
	Nonce 		int
}

func CreateBlock(data string, prevHash []byte) *Block {
	// block constructor 
	block := &Block{[] byte{}, []byte(data), prevHash, 0}

	// create proof of work by passing in block
	pow := NewProof(block)

	// executing run function on that proof of work
	nonce, hash := pow.Run() 

	block.Hash = hash[:]
	block.Nonce = nonce
	
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
} 

func Gensis() *Block {
	return CreateBlock("gensis", []byte{})
}

func InitBlockChain() *BlockChain {
	// reference to blockchain, and create array of blocks with call to Genesis function
	return &BlockChain{[]*Block{Gensis()}}
}