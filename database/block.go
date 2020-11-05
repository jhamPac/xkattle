package database

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// Hash represents the cryptographic has of a block
type Hash [32]byte

// MarshalText encodes a hash to a string
func (h Hash) MarshalText() ([]byte, error) {
	return []byte(hex.EncodeToString(h[:])), nil
}

// UnmarshalText decodes the hash value into the actual number
func (h *Hash) UnmarshalText(data []byte) error {
	_, err := hex.Decode(h[:], data)
	return err
}

// dbFS represents the meta portion of a hash and block in the db file
type dbFS struct {
	Key   Hash  `json:"hash"`
	Value Block `json:"block"`
}

// Block represents a block on the chain with a group of txs
type Block struct {
	Header BlockHeader `json:"header"`
	TXs    []Tx        `json:"payload"`
}

// BlockHeader is meta data about a block
type BlockHeader struct {
	Parent Hash   `json:"parent"`
	Time   uint64 `json:"time"`
}

// NewBlock instantiates and returns a Block
func NewBlock(parent Hash, time uint64, txs []Tx) Block {
	return Block{BlockHeader{parent, time}, txs}
}

// Hash creates a cryptographic hash of a block
func (b Block) Hash() (Hash, error) {
	blockJSON, err := json.Marshal(b)
	if err != nil {
		return Hash{}, err
	}
	return sha256.Sum256(blockJSON), nil
}
