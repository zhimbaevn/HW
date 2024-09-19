package cmd

import (
	"fmt"
	"log"
	"orders/internal/app/command/neworder"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "take a new order from a courier",
	Long: `Take a new order from a courier
As input, the command takes three required parameters: the order ID, the client ID, and the storage period
Example: orders new --client=585 --order=67845 --storage=2024-11-05`,
	RunE: func(cmd *cobra.Command, args []string) error {
		const op = "new.newCmd.Run"

		orderId, err := cmd.Flags().GetInt("order")
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		clietnId, err := cmd.Flags().GetInt("client")
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		storagePeriod, err := cmd.Flags().GetString("storage")
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		err = neworder.New(Storage, orderId, clietnId, storagePeriod)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		log.Println("Order has been recorded")
		return nil
	},
}
