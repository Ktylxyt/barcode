package cmd

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/spf13/cobra"
)

var BarCode = &cobra.Command{
	Use:   "barcode",
	Short: "generate new bar code and qr code",
	Run:   Generate,
}

func Generate(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("barcode: no file specified")
		return
	}

	fmt.Println("axaxaxaxaxaxa")

	qrCode, _ := qr.Encode(args[0], qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	file, _ := os.Create(args[1])
	defer file.Close()

	png.Encode(file, qrCode)
}

func init() {
	rootCmd.AddCommand(BarCode)
}
