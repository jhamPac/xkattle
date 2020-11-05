package database

type Hash [32]byte

type Block struct {
	Header BlockHeader `json:"header"`
	TXs    []Tx        `json:"payload"`
}

type BlockHeader struct {
	Parent Hash
	Time   uint64
}
