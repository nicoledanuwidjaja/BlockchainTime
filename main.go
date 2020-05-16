package main

import (
  "fmt"
  "bytes"
  "crypto/sha256"
)

type BlockChain struct {
  blocks []*Block
}

type Block struct {
  Hash []byte
  Data []byte
  PrevHash []byte
}

func (chain *BlockChain) AddBlock(data string) {
  prevBlock := chain.blocks[len(chain.blocks) - 1]
  newBlock := CreateBlock(data, prevBlock.Hash)
  chain.blocks = append(chain.blocks, newBlock)
}

func (b*Block) GenerateHash() {
  info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
  hash := sha256.Sum256(info)
  b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
  block := &Block{[]byte{}, []byte(data), prevHash}i
  block.DeriveHash()

  return block
}


func main() {
  fmt.Println(quote.Hello())
}
