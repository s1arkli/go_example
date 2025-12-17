package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd")
	},
}

func Execute() {
	rootCmd.Execute() //命令主体
}
