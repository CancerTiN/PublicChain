package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	// 1. Block height
	Height int64
	// 2. HASH of the previous block
	PrevBlockHash []byte
	// 3. Transaction data
	Data []byte
	// 4. Timestamp
	Timestamp int64
	// 5. Hash
	Hash []byte
}

func (block *Block) SetHash() {
	// 1. Convert block height to byte slice
	heightBytes := Int64ToHex(block.Timestamp)
	// 2. Convert timestamps to byte slices
	timeString := strconv.FormatInt(block.Timestamp, 2)
	timeBytes := []byte(timeString)
	// 3. Stitch all attributes
	blockBytes := bytes.Join([][]byte{heightBytes, block.PrevBlockHash, block.Data, timeBytes, block.Hash}, []byte{})
	// 4. Generate a hash of the current block
	byteArray := sha256.Sum256(blockBytes)
	block.Hash = byteArray[:]
}

// Create new block
func NewBlock(data string, height int64, preBlockHash []byte) (block *Block) {
	// Create block
	block = &Block{
		Height:height,
		PrevBlockHash:preBlockHash,
		Data:[]byte(data),
		Timestamp:time.Now().Unix(),
	}
	// Set hash
	block.SetHash()
	return
}
