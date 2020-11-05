package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

const major = "0"
const minor = "4"
const patch = "0"
const desc = "Blockchain in Go"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays xpay version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s.%s.%s-beta %s", major, minor, patch, desc)
	},
}
