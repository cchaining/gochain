/**
gc_core 패키지 :

block, blockchain, chain_iterator를 담고있는 패키지이다.
chain과 관련된 go 파일을 담고있다. ethereum으로 비교하자면 core 패키지와 같다.

 */
package gc_core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"strconv"
	"time"

)

// Block represents a block in the blockchain
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}


func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}


func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}


// 블록 직렬화
func (b *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)

	return result.Bytes()
}


// 블록 역직렬화
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)

	return *block
}


func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}