package cmd

import (
	"fmt"
	"log/slog"
	"orders/internal/storage/jsondb"

	"github.com/spf13/cobra"
)

var Storage jsondb.Storage
var Log *slog.Logger

var rootCmd = &cobra.Command{
	Use:   "orders",
	Short: "managing orders at the pickup point",
	Long:  ``,
}

func Execute(s *jsondb.Storage) error {
	const op = "root.Execute"
	Storage = *s

	rootCmd.CompletionOptions.DisableDefaultCmd = true // disable default completion

	err := rootCmd.Execute()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func init() {

	// add subcommands new
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().IntP("order", "o", 0, "order id")
	newCmd.MarkFlagRequired("order")

	newCmd.Flags().IntP("client", "c", 0, "client id")
	newCmd.MarkFlagRequired("client")

	newCmd.Flags().StringP("storage", "s", "2006-01-02", "storage priod date in YYYY-MM-DD format")
	newCmd.MarkFlagRequired("storage")

	// add subcommands del
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntP("id", "i", 0, "order id")

	// add subcommands get
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().IntP("client", "c", 0, "client ID")
	getCmd.MarkFlagRequired("client")

	getCmd.Flags().BoolP("unissued", "u", false, "return list of orders that are still stored")
	getCmd.Flags().IntP("number", "n", 0, "return list with last <number> orders")

	// add subcommands issue
	rootCmd.AddCommand(issueCmd)
	issueCmd.Flags().IntSliceP("id", "i", []int{}, "Один или несколько заказов")

	// add subcommands refund
	rootCmd.AddCommand(refundCmd)
	refundCmd.Flags().IntP("order", "o", 0, "order id")
	refundCmd.Flags().IntP("client", "c", 0, "client id")

	// add subcommands refls
	rootCmd.AddCommand(refLsCmd)
	refLsCmd.Flags().IntP("lines", "l", 5, "number of lines per page")
}
