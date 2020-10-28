package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "xpay",
		Short: "xKattle blockchain CLI.",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	rootCmd.AddCommand(versionCmd)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
