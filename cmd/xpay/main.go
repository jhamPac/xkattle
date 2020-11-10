package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const flagDataDir = "datadir"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "xpay",
		Short: "xKattle blockchain CLI.",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(balancesCmd())
	rootCmd.AddCommand(txCmd())

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addDefaultRequiredFlags(cmd *cobra.Command) {
	cmd.Flags().String(flagDataDir, "", "Absolute path to the node data dir where the DB will/is stored")
	cmd.MarkFlagRequired(flagDataDir)
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
