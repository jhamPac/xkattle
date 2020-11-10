package main

import (
	"fmt"
	"os"

	"github.com/jhampac/xkattle/database"
	"github.com/spf13/cobra"
)

func balancesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balances",
		Short: "Interact with the balances (list...)",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	// xpay balances list ...
	cmd.AddCommand(balancesListCmd())
	return cmd
}

func balancesListCmd() *cobra.Command {
	var bcmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all balances",
		Run: func(cmd *cobra.Command, args []string) {
			dataDir, _ := cmd.Flags().GetString(flagDataDir)
			state, err := database.NewStateFromDisk(dataDir)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			defer state.Close()

			fmt.Printf("Account balances at %x:\n", state.LatestBlockHash())
			fmt.Println("_________________")
			fmt.Println("")

			for account, balance := range state.Balances {
				fmt.Println(fmt.Sprintf("%s: %d", account, balance))
			}
		},
	}

	addDefaultRequiredFlags(bcmd)

	return bcmd
}
