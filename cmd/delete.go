package cmd

import (
	"fmt"
	"log"
	"orders/internal/app/command/deleteorder"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "del",
	Short: "return order to courier",
	Long: `Return order to courier by order ID
Example: orders del --id=1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		const op = "delete.deleteCmd.Run"

		orderId, err := cmd.Flags().GetInt("id")
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		err = deleteorder.DeleteOrder(Storage, orderId)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		log.Printf("Order %d returned\n", orderId)

		return nil
	},
}
