package getorder

import (
	"fmt"
	"log"
	"orders/internal/app/order"
	"orders/internal/storage/jsondb"
	"sort"
	"time"
)

const recLayout = "2006-01-02 15:04:05.999999999"

func GetUnissued(s jsondb.Storage, cid int) (*map[int]order.Order, error) {
	const op = "command.getorder.GetUnissued"

	dataPtr, err := s.GetData()
	if err != nil {
		return &map[int]order.Order{}, fmt.Errorf("%s: %w", op, err)
	}

	data := *dataPtr
	orders := make(map[int]order.Order)

	for _, order := range data {
		if order.ClientID == cid && order.IssueDate == "" {
			orders[order.OrderId] = order
		}
	}

	return &orders, nil
}

func GetOrders(s jsondb.Storage, cid int) ([]order.Order, error) {
	const op = "command.getorder.GetOrders"

	dataPtr, err := s.GetData()
	if err != nil {
		return []order.Order{}, fmt.Errorf("%s: %w", op, err)
	}

	data := *dataPtr
	orders := []order.Order{}

	for _, order := range data {
		if order.ClientID == cid {
			orders = append(orders, order)
		}
	}

	sort.Slice(orders, func(i, j int) bool {
		time_i, err := time.Parse(recLayout, orders[i].RecTime)
		if err != nil {
			log.Printf("%s: %s", op, err)
			return false
		}

		time_j, err := time.Parse(recLayout, orders[j].RecTime)
		if err != nil {
			log.Printf("%s: %s", op, err)
			return false
		}

		return time_i.Before(time_j)
	})

	return orders, nil

}
