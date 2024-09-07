/*
Copyright © 2024 Nikolai Zhimbaev>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"orders/internal/storage/jsondb"

	"github.com/spf13/cobra"
)

var Storage jsondb.Storage
var Log *slog.Logger

var rootCmd = &cobra.Command{
	Use:   "orders",
	Short: "Утилита для управлением заказами пункта выдачи",
	Long: `Утилита для управлением заказами пункта выдачи.
На данный момент есть возможность:
Принять заказ от курьера - new,
Вернуть заказ курьеру - del,
Выдать заказ клиенту - issue
Вывести список заказов клиента - get
Принять возврат от клиента - refund
Получить список возвратов - refls`,
}

func Execute(s *jsondb.Storage, log *slog.Logger) error {
	const op = "root.Execute"
	Storage = *s
	Log = log

	err := rootCmd.Execute()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func init() {

}
