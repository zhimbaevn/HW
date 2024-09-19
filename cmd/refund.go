package cmd

import (
	"fmt"
	"log"
	"orders/internal/app/command/refundorder"

	"github.com/spf13/cobra"
)

var refundCmd = &cobra.Command{
	Use:   "refund",
	Short: "accept a refund from a client.",
	Long: `Accept a refund from a client.
As input, the command takes two required parameters: the order ID and the user ID.
Example: orders refund --order 1 --client 1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		const op = "refund.refuncCmd.Run"
		orderId, err := cmd.Flags().GetInt("order")
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		clietnId, err := cmd.Flags().GetInt("client")
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		err = refundorder.RefundOrder(Storage, orderId, clietnId)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		log.Printf("Client %d refunded order %d\n", clietnId, orderId)
		return nil
	},
}
