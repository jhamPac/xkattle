package main

import (
	"fmt"
	"os"

	"github.com/jhampac/xkattle/database"
	"github.com/spf13/cobra"
)

const (
	flagFrom  = "from"
	flagTo    = "to"
	flagValue = "value"
	flagData  = "data"
)

func txCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tx",
		Short: "Interact with txs (add...)",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.AddCommand(txAddCmd())
	return cmd
}

func txAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Adds new TX to database",
		Run: func(cmd *cobra.Command, args []string) {
			from, _ := cmd.Flags().GetString(flagFrom)
			to, _ := cmd.Flags().GetString(flagTo)
			value, _ := cmd.Flags().GetUint(flagValue)
			data, _ := cmd.Flags().GetString(flagData)

			tx := database.NewTx(database.NewAccount(from), database.NewAccount(to), value, data)

			state, err := database.NewStateFromDisk()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			defer state.Close()

			err = state.AddTx(tx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			_, err = state.Persist()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			fmt.Println("TX successfully added to the ledger")
		},
	}

	cmd.Flags().String(flagFrom, "", "Payer account")
	cmd.MarkFlagRequired(flagFrom)

	cmd.Flags().String(flagTo, "", "Payee account")
	cmd.MarkFlagRequired(flagTo)

	cmd.Flags().Uint(flagValue, 0, "Token amount to send")
	cmd.MarkFlagRequired(flagValue)

	cmd.Flags().String(flagData, "", "Possible values: 'reward'")

	return cmd
}
