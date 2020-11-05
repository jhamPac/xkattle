package database

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// State is the main business logic of the ledger
type State struct {
	Balances  map[Account]uint
	txMempool []Tx

	dbFile          *os.File
	latestBlockHash Hash
}

// NewStateFromDisk starts the ledger from the genesis
func NewStateFromDisk() (*State, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	gen, err := loadGenesis(filepath.Join(cwd, "database", "genesis.json"))
	if err != nil {
		return nil, err
	}

	balances := make(map[Account]uint)
	for account, balance := range gen.Balances {
		balances[account] = balance
	}

	f, err := os.OpenFile(filepath.Join(cwd, "database", "block.db"), os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)

	state := &State{balances, make([]Tx, 0), f, Hash{}}

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		metaJSON := scanner.Bytes()
		var meta BlockMeta
		err = json.Unmarshal(metaJSON, &meta)
		if err != nil {
			return nil, err
		}

		err = state.applyBlock(meta.Value)
		if err != nil {
			return nil, err
		}

		state.latestBlockHash = meta.Key
	}

	return state, nil
}

// LatestBlockHash returns the current hash of the tx.db file
func (s *State) LatestBlockHash() Hash {
	return s.latestBlockHash
}

// AddBlock iterates over each tx in a block and calls Addtx
func (s *State) AddBlock(b Block) error {
	for _, tx := range b.TXs {
		if err := s.AddTx(tx); err != nil {
			return err
		}
	}
	return nil
}

// AddTx adds a tx to the state mempool
func (s *State) AddTx(tx Tx) error {
	if err := s.apply(tx); err != nil {
		return err
	}
	s.txMempool = append(s.txMempool, tx)
	return nil
}

// Persist mempool to disk
func (s *State) Persist() (Hash, error) {
	block := NewBlock(s.latestBlockHash, uint64(time.Now().Unix()), s.txMempool)
	blockHash, err := block.Hash()
	if err != nil {
		return Hash{}, err
	}

	meta := BlockMeta{blockHash, block}
	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return Hash{}, err
	}

	fmt.Printf("Persisting new Block to disk:\n")
	fmt.Printf("\t%s\n", metaJSON)

	if _, err = s.dbFile.Write(append(metaJSON, '\n')); err != nil {
		return Hash{}, err
	}

	s.latestBlockHash = blockHash

	// reset mempool
	s.txMempool = []Tx{}
	return s.latestBlockHash, nil
}

// Close references to the file
func (s *State) Close() {
	s.dbFile.Close()
}

func (s *State) applyBlock(b Block) error {
	for _, tx := range b.TXs {
		if err := s.apply(tx); err != nil {
			return err
		}
	}
	return nil
}

func (s *State) apply(tx Tx) error {
	if tx.IsReward() {
		s.Balances[tx.To] += tx.Value
		return nil
	}

	if s.Balances[tx.From] < tx.Value {
		return fmt.Errorf("insufficient balance")
	}

	s.Balances[tx.From] -= tx.Value
	s.Balances[tx.To] += tx.Value

	return nil
}
