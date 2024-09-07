package cmd

import (
	"fmt"
	"orders/internal/app/command/neworder"
	"orders/internal/lib/logger/sl"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Принять заказ от курьера",
	Long: `Приём заказ от курьера.
На вход подаётся 3 обязательных параметра id заказа, id клиента, срок хранения.
Пример: orders new --client=585 --order=67845 --storage=2024-11-05`,
	Run: func(cmd *cobra.Command, args []string) {
		const op = "new.newCmd.Run"

		orderId, err := cmd.Flags().GetInt("order")
		if err != nil {
			Log.Error("can't parse --order", sl.Err(fmt.Errorf("%s: %w", op, err)))
			return
		}
		clietnId, err := cmd.Flags().GetInt("client")
		if err != nil {
			Log.Error("can't parse --client", sl.Err(fmt.Errorf("%s: %w", op, err)))
			return
		}
		storagePeriod, err := cmd.Flags().GetString("storage")
		if err != nil {
			Log.Error("can't parse --storage", sl.Err(fmt.Errorf("%s: %w", op, err)))
			return
		}

		err = neworder.New(Storage, orderId, clietnId, storagePeriod)
		if err != nil {
			Log.Error("can't save to storage", sl.Err(fmt.Errorf("%s: %w", op, err)))
			return
		}

		Log.Info("Order has been recorded")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().IntP("order", "o", 0, "order id")
	newCmd.MarkFlagRequired("order")

	newCmd.Flags().IntP("client", "c", 0, "client id")
	newCmd.MarkFlagRequired("client")

	newCmd.Flags().StringP("storage", "s", "2006-01-02", "storage priod date in YYYY-MM-DD format")
	newCmd.MarkFlagRequired("storage")
}
