package refundlist

import (
	"fmt"
	"orders/internal/app/order"
	"orders/internal/storage/jsondb"
)

func GetRefund(s jsondb.Storage) ([]order.Order, error) {
	const op = "command.refundlist.GetRefund"

	dataPtr, err := s.GetData()
	if err != nil {
		return []order.Order{}, fmt.Errorf("%s: %w", op, err)
	}

	data := *dataPtr
	orders := []order.Order{}

	for _, order := range data {
		if order.RefundDate != "" {
			orders = append(orders, order)
		}
	}

	return orders, nil
}
