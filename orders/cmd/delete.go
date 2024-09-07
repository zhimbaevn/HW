package cmd

import (
	"fmt"
	"orders/internal/app/command/deleteorder"
	"orders/internal/lib/logger/sl"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "del",
	Short: "Вернуть заказ курьеру",
	Long: `Возврат заказа курьеру.
Пример: orders del --id=585`,
	Run: func(cmd *cobra.Command, args []string) {
		const op = "delete.deleteCmd.Run"

		orderId, err := cmd.Flags().GetInt("id")
		if err != nil {
			Log.Error("failed to parse order id", sl.Err(fmt.Errorf("%s: %w", op, err)))
			return
		}

		err = deleteorder.DeleteOrder(Storage, orderId)
		if err != nil {
			Log.Error("failed to return order", sl.Err(fmt.Errorf("%s: %w", op, err)))
			return
		}

		Log.Info("Order has been removed from the database")

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntP("id", "i", 0, "order id")
}
