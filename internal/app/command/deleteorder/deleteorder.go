package deleteorder

import (
	"errors"
	"fmt"
	"orders/internal/storage"
	"orders/internal/storage/jsondb"
	"time"
)

func DeleteOrder(s jsondb.Storage, orderId int) error {
	const layout = "2006-01-02"
	const op = "command.deleteorder.deleteOrder"

	dataPtr, err := s.GetData()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	data := *dataPtr

	v, ok := data[orderId]
	if !ok {
		return fmt.Errorf("%s: %w", op, storage.ErrOrderNotFound)
	}

	ErrStoragePeriodNotExpired := errors.New("order storage period has not yet expired")
	ErrOrderIssued := errors.New("order has been issued to client")

	orderDate, err := time.Parse(layout, v.StoragePeriod)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if orderDate.After(time.Now()) {
		return fmt.Errorf("%s: %w", op, ErrStoragePeriodNotExpired)
	}

	if v.IssueDate != "" {
		return fmt.Errorf("%s: %w", op, ErrOrderIssued)
	}

	delete(data, orderId)

	err = s.SaveData(&data)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil

}
