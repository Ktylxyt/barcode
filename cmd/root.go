package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "bar",
	Short: "generate new barcode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate new barcode")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

