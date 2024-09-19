package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"orders/internal/app/command/refundlist"
	"orders/internal/app/order"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var refLsCmd = &cobra.Command{
	Use:   "refls",
	Short: "display a list of refunds",
	Long: `Display a list of refunds.
Navigation is performed using the 'n' and 'p' keys.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		const op = "refundlist.refLsCmd.Run"

		pageSize, err := cmd.Flags().GetInt("lines")
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		data, err := refundlist.GetRefund(Storage)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		page := 0
		reader := bufio.NewReader(os.Stdin)

		if len(data) == 0 {
			return fmt.Errorf("%s: %w", op, errors.New("no refunds found"))
		}

		for {
			clearConsole()
			printPage(data, page, pageSize)

			fmt.Printf("\nСтраница %d из %d. Нажмите 'n' для следующей страницы, 'p' для предыдущей страницы, 'q' для выхода.\n", page+1, (len(data)+pageSize-1)/pageSize)
			input, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
			input = strings.TrimSpace(input)

			switch input {
			case "n":
				if (page+1)*pageSize < len(data) {
					page++
				}
			case "p":
				if page-1 >= 0 {
					page--
				}
			case "q":
				return nil
			default:
				fmt.Println("Неизвестная команда")
				fmt.Printf("Список команд:\nn - следующая страница\np - предыдущая страница\nq - выход\n\n")
			}
		}

	},
}

func clearConsole() {
	cmd := exec.Command("clear") // Для Linux и Mac
	if err := cmd.Run(); err != nil {
		cmd = exec.Command("cls") // Для Windows
		cmd.Run()
	}
}

func printPage(data []order.Order, page int, pageSize int) {
	start := page * pageSize
	end := start + pageSize

	if start >= len(data) {
		fmt.Println("Нет данных для отображения.")
		return
	}

	if end > len(data) {
		end = len(data)
	}

	for _, item := range data[start:end] {
		fmt.Println(item)
	}
}
