package cmd

import (
	"errors"
	"fmt"
	"log"
	"orders/internal/app/command/issueorder"

	"github.com/spf13/cobra"
)

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "issuing orders to a client",
	Long: `Issue orders to a client.
As input, the command takes a slice of order IDs
example: orders issue -i 55,22,33`,
	RunE: func(cmd *cobra.Command, args []string) error {
		const op = "issue.issueCmd.Run"

		orders, err := cmd.Flags().GetIntSlice("id")
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		if len(orders) < 1 {
			return fmt.Errorf("%s: %w", op, errors.New("length slice of orders is 0"))
		}
		err = issueorder.IssueOrder(Storage, orders)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		log.Println("All existing orders have been issued")

		return nil
	},
}
