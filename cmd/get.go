package cmd

import (
	"errors"
	"fmt"
	"log"
	"orders/internal/app/command/getorder"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get client order list",
	Long: `Return order list by client ID
Example: orders get --client 1 -n 3
	 orders get --client 1 -u`,

	RunE: func(cmd *cobra.Command, args []string) error {
		const op = "get.getCmd.Run"

		clientId, err := cmd.Flags().GetInt("client")
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		unissued, err := cmd.Flags().GetBool("unissued")
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		if !unissued {

			count, err := cmd.Flags().GetInt("number")
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
			if count <= 0 {
				return fmt.Errorf("%s: %w", op, errors.New("no arguments were passed or --number is negative"))
			}

			//Вывод n заказов
			orders, err := getorder.GetOrders(Storage, clientId)
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}

			c := min(len(orders), count)

			for i := 1; i <= c; i++ {
				fmt.Printf("%d: %s, recording time: %s\n", i, orders[len(orders)-i], orders[len(orders)-i].RecTime)
			}

			return nil

		}

		//Если нужны невыданные заказы
		ordersPtr, err := getorder.GetUnissued(Storage, clientId)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		orders := *ordersPtr

		if len(orders) == 0 {
			log.Println("The client with this ID does not have any orders at our pick-up point.")
			return nil
		}

		fmt.Printf("%d orders found in database:\n", len(orders))

		for _, order := range orders {
			fmt.Printf("Order ID: %d | Storage period: %s\n", order.OrderId, order.StoragePeriod)
		}

		log.Println("End of order list")
		return nil
	},
}
