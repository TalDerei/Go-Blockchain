package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

const Difficulty = 18

type ProofOfWork struct {
	Block *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)

	// left shift
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow 
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

func ToHex(num int64) []byte {
	// create buffer
	buff := new(bytes.Buffer) 

	// binary Write to decode number into bytes 
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce) 
		hash := sha256.Sum256(data)

		fmt.Printf("\r%x", hash)

		// convert hash to big integer
		intHash.SetBytes(hash[:])

		// compare target with big int hash
		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
		fmt.Println()
	}
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var initHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	initHash.SetBytes(hash[:])

	return initHash.Cmp(pow.Target) == -1
}
