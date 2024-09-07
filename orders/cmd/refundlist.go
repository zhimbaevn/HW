package cmd

import (
	"bufio"
	"fmt"
	"orders/internal/app/command/refundlist"
	"orders/internal/app/order"
	"orders/internal/lib/logger/sl"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var refLsCmd = &cobra.Command{
	Use:   "refls",
	Short: "Вывести список возвратов",
	Long: `Постраничный список возвратов.
Для навигации используйте n и p или вручную введите номер требуемой страницы`,
	Run: func(cmd *cobra.Command, args []string) {
		const op = "refundlist.refLsCmd.Run"

		pageSize, err := cmd.Flags().GetInt("lines")
		if err != nil {
			Log.Error("can't parse --lines", sl.Err(fmt.Errorf("%s: %w", op, err)))
			return
		}

		data, err := refundlist.GetRefund(Storage)
		if err != nil {
			Log.Error("can't get data", sl.Err(fmt.Errorf("%s: %w", op, err)))
			return
		}

		page := 0
		reader := bufio.NewReader(os.Stdin)

		if len(data) == 0 {
			Log.Info("No refunds found")
			return
		}

		for {
			clearConsole()
			printPage(data, page, pageSize)

			fmt.Printf("\nСтраница %d из %d. Нажмите 'n' для следующей страницы, 'p' для предыдущей страницы, 'q' для выхода.\n", page+1, (len(data)+pageSize-1)/pageSize)
			input, err := reader.ReadString('\n')
			if err != nil {
				Log.Error("can't ReadString from Stdin", sl.Err(fmt.Errorf("%s: %w", op, err)))
				return
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
				return
			default:
				fmt.Println("Неизвестная команда")
				fmt.Printf("Список команд:\nn - следующая страница\np - предыдущая страница\nq - выход\n\n")
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(refLsCmd)
	refLsCmd.Flags().IntP("lines", "l", 5, "number of lines per page")
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
