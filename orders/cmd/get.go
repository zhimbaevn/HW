package cmd

import (
	"fmt"
	"orders/internal/app/command/getorder"
	"orders/internal/lib/logger/sl"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Список заказов клиента",
	Long: `Команда даёт возможность получить список последних n заказов или только ожидающие выдачи заказы.
Параметр --client - id клиента, заказы которого нужно сформировать. *Обязательный параметр
Если передан параметр -n, то вернётся список с последними n заказами клиента
Если передан параметр -u, то вернётся список заказов, хранящихся на ПВЗ`,

	Run: func(cmd *cobra.Command, args []string) {
		const op = "get.getCmd.Run"

		clientId, err := cmd.Flags().GetInt("client")
		if err != nil {
			Log.Error("can't parse --client", sl.Err(fmt.Errorf("%s: %w", op, err)))
			return
		}

		unissued, err := cmd.Flags().GetBool("unissued")
		if err != nil {
			Log.Error("can't parse --unissued", sl.Err(fmt.Errorf("%s: %w", op, err)))
			return
		}

		if !unissued {

			count, err := cmd.Flags().GetInt("number")
			if err != nil {
				Log.Error("can't parse --number", sl.Err(fmt.Errorf("%s: %w", op, err)))
				return
			}
			if count <= 0 {
				Log.Info("No arguments were passed or --number is negative")
				return
			}

			//Вывод n заказов
			orders, err := getorder.GetOrders(Storage, clientId)
			if err != nil {
				Log.Error("can't get orders by numbers", sl.Err(fmt.Errorf("%s: %w", op, err)))
				return
			}

			c := min(len(orders), count)

			for i := 1; i <= c; i++ {
				fmt.Printf("%d: %s, recording time: %s\n", i, orders[len(orders)-i], orders[len(orders)-i].RecTime)
			}

			return

		}

		//Если нужны невыданные заказы
		ordersPtr, err := getorder.GetUnissued(Storage, clientId)
		if err != nil {
			Log.Error("can't  GetUnissued", sl.Err(fmt.Errorf("%s: %w", op, err)))
			return
		}
		orders := *ordersPtr

		if len(orders) == 0 {
			Log.Info("The client with this ID does not have any orders at our pick-up point.")
			return
		}

		fmt.Printf("%d orders found in database:\n", len(orders))

		for _, order := range orders {
			fmt.Printf("Order ID: %d | Storage period: %s\n", order.OrderId, order.StoragePeriod)
		}

		Log.Info("End of order list")

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().IntP("client", "c", 0, "id клиента *обязательный флаг")
	getCmd.MarkFlagRequired("client")

	getCmd.Flags().BoolP("unissued", "u", false, "Если передан параметр --unissued, то вернётся список заказов, хранящихся на ПВЗ")
	getCmd.Flags().IntP("number", "n", 0, "Если передан параметр -number, то вернётся список с последними n заказами клиента")
}
