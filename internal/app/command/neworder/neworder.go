package neworder

import (
	"errors"
	"fmt"
	"orders/internal/app/order"
	"orders/internal/storage"
	"orders/internal/storage/jsondb"
	"time"
)

const (
	layout    = "2006-01-02"
	recLayout = "2006-01-02 15:04:05.999999999"
)

func New(s jsondb.Storage, orderId, clientId int, storagePeriod string) error {
	const op = "cmd.new.New"

	//Проверка на сущетсвования такого товра
	dataPtr, err := s.GetData()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	data := *dataPtr

	_, ok := data[orderId]
	if ok {
		return fmt.Errorf("%s: %w", op, storage.ErrOrderExist)
	}

	// Проверка на срок хранения
	ErrStoragePeriodExpired := errors.New("order storage period has already expired")

	orderDate, err := time.Parse(layout, storagePeriod)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	today := time.Now()

	if orderDate.Before(today) {
		return fmt.Errorf("%s: %w", op, ErrStoragePeriodExpired)
	}

	// SAVE
	t := time.Now().Format(recLayout)

	data[orderId] = order.Order{OrderId: orderId, ClientID: clientId, StoragePeriod: storagePeriod, RecTime: t}

	err = s.SaveData(&data)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
