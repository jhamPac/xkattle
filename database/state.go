package database

import (
	"os"
	"path/filepath"
)

// State is the main business logic of the ledger
type State struct {
	Balances  map[Account]uint
	txMempool []Tx

	dbFile *os.File
}

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
}
