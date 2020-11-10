package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jhampac/xkattle/database"
)

func main() {
	cwd, _ := os.Getwd()
	state, err := database.NewStateFromDisk(cwd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer state.Close()

	block0 := database.NewBlock(
		database.Hash{},
		uint64(time.Now().Unix()),
		[]database.Tx{
			database.NewTx("alice", "alice", 3, ""),
			database.NewTx("alice", "alice", 700, "reward"),
		},
	)

	state.AddBlock(block0)
	block0hash, _ := state.Persist()

	block1 := database.NewBlock(
		block0hash,
		uint64(time.Now().Unix()),
		[]database.Tx{
			database.NewTx("alice", "bob", 2000, ""),
			database.NewTx("alice", "alice", 100, "reward"),
			database.NewTx("bob", "alice", 1, ""),
			database.NewTx("alice", "eve", 10000, ""),
			database.NewTx("alice", "bob", 10000, ""),
			database.NewTx("alice", "charlie", 1000, ""),
			database.NewTx("alice", "charlie", 1000, ""),
			database.NewTx("alice", "alice", 1000, "reward"),
			database.NewTx("charlie", "charlie", 10000, "reward"),
			database.NewTx("charlie", "charlie", 10000, "reward"),
		},
	)

	state.AddBlock(block1)
	state.Persist()
}
