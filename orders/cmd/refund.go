package cmd

import (
	"fmt"
	"orders/internal/app/command/refundorder"
	"orders/internal/lib/logger/sl"

	"github.com/spf13/cobra"
)

var refundCmd = &cobra.Command{
	Use:   "refund",
	Short: "Принять возврат от клиента",
	Long: `Возврат заказа от клиента.
На вход принимается два обязательных параметра id заказа и id пользователя`,
	Run: func(cmd *cobra.Command, args []string) {
		const op = "refund.refuncCmd.Run"
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

		err = refundorder.RefundOrder(Storage, orderId, clietnId)
		if err != nil {
			Log.Error("can't refund order", sl.Err(fmt.Errorf("%s: %w", op, err)))
			return
		}

		s := fmt.Sprintf("Client %d refunded order %d", clietnId, orderId)
		Log.Info(s)

	},
}

func init() {
	rootCmd.AddCommand(refundCmd)
	refundCmd.Flags().IntP("order", "o", 0, "order id")
	refundCmd.Flags().IntP("client", "c", 0, "client id")
}
