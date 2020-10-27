package database

// Account is a stirng with xkattle specific functionality to represent an account
type Account string

// NewAccount creates and returns an Account
func NewAccount(value string) Account {
	return Account(value)
}

// Tx represents a transaction
type Tx struct {
	From  Account `json:"from"`
	To    Account `json:"to"`
	Value uint    `json:"value"`
	Data  string  `json:"data"`
}

// NewTx creates and returns a transaction
func NewTx(from Account, to Account, value uint, data string) Tx {
	return Tx{from, to, value, data}
}

// IsReward checks weather the transaction is an reward or not
func (t Tx) IsReward() bool {
	return t.Data == "reward"
}
