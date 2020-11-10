package database

import (
	"encoding/json"
	"io/ioutil"
)

var genesisJSON = `
{
	"genesis_time": "2020-11-11T00:00:00.000000000Z",
	"chain_id": "xkattle-ledger",
	"balances": {
		"alice": 1000000
  	}
}`

type genesis struct {
	Balances map[Account]uint `json:"balances"`
}

func loadGenesis(path string) (genesis, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return genesis{}, err
	}

	var g genesis
	err = json.Unmarshal(content, &g)
	if err != nil {
		return genesis{}, nil
	}
	return g, nil
}

func writeGenesisToDisk(path string) error {
	return ioutil.WriteFile(path, []byte(genesisJSON), 0644)
}
