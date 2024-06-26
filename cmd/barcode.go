package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var barcode = &cobra.Command{
	Use: "barcode",
	Short: "generate new bar code and qr code",
	Run: Generate,
}

func Generate(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("barcode: no file specified")
		return
	}
}

func init() {
	rootCmd.AddCommand(barcode)
}