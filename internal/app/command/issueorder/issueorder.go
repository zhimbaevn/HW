package issueorder

import (
	"fmt"
	"log"
	"orders/internal/storage"
	"orders/internal/storage/jsondb"
	"time"
)

func IssueOrder(s jsondb.Storage, orders []int) error {
	const (
		op     = "command.issueOrder.IssueOrder"
		layout = "2006-01-02"
	)

	dataPtr, err := s.GetData()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	//TODO Все ID заказов должны принадлежать только одному клиенту
	clients := make(map[int][]int) // map[client id] orderid
	today := time.Now().Format(layout)

	for _, o := range orders {
		data := *dataPtr
		order, ok := data[o]

		if !ok {
			log.Println(op, "INFO: Order №", o, "does not exist")
			continue
		}
		clients[order.ClientID] = append(clients[order.ClientID], order.OrderId)
		// Если больше одного клиента
		if len(clients) > 1 {
			return fmt.Errorf("%s: %w", op, storage.ErrAttemptIssueFewClients)
		}

		order.IssueDate = today
		data[o] = order

	}

	if len(clients) == 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrNoOneIssed)
	}

	err = s.SaveData(dataPtr)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
