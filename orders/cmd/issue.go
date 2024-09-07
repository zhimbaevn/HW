package cmd

import (
	"errors"
	"fmt"
	"orders/internal/app/command/issueorder"
	"orders/internal/lib/logger/sl"
	"orders/internal/storage"

	"github.com/spf13/cobra"
)

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Выдать заказы клиенту",
	Long: `Выдача заказов клиенту. На вход подаётся слайс из id заказов
Если заказы в данный момент на ПВЗ, не истёк срок хранения и все они принадлежат одному клинету
то в хранилище появятся отметка об их выдаче сегодняшней датой.
Пример: orders issue -i 55,22,33`,
	Run: func(cmd *cobra.Command, args []string) {
		const op = "issue.issueCmd.Run"

		orders, err := cmd.Flags().GetIntSlice("id")
		if err != nil {
			Log.Error("failed to parse slice of orders id", sl.Err(fmt.Errorf("%s: %w", op, err)))
			return
		}
		if len(orders) < 1 {
			Log.Error("length slice of orders is 0")
			return
		}
		err = issueorder.IssueOrder(Storage, orders)
		if err != nil {
			if errors.Is(err, storage.ErrNoOneIssed) {
				Log.Info("No orders were issued")
				return
			}
			Log.Error("failed to issue order", sl.Err(fmt.Errorf("%s: %w", op, err)))
			return
		}

		Log.Info("All existing orders have been issued")

	},
}

func init() {
	rootCmd.AddCommand(issueCmd)
	issueCmd.Flags().IntSliceP("id", "i", []int{}, "Один или несколько заказов")
}
