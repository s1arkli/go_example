package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(firstCmd)
	firstCmd.AddCommand(secondCmd)
}

var firstCmd = &cobra.Command{
	Use: "first",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("first cmd")
	},
}

var secondCmd = &cobra.Command{
	Use: "second",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("second cmd")
	},
}
