package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var xpayCmd = &cobra.Command{
		Use:   "xpay",
		Short: "xKattle blockchain CLI",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello")
		},
	}

	err := xpayCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
