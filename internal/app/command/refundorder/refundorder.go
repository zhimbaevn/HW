package refundorder

import (
	"fmt"
	"orders/internal/storage"
	"orders/internal/storage/jsondb"
	"time"
)

const layout = "2006-01-02"

func RefundOrder(s jsondb.Storage, orderId, clientId int) error {
	const op = "cmd.refundorder.RefundOrder"

	dataPtr, err := s.GetData()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	data := *dataPtr

	_, ok := data[orderId]
	if !ok {
		return fmt.Errorf("%s: %w", op, storage.ErrOrderNotFound)
	}

	if data[orderId].RefundDate != "" {
		return fmt.Errorf("%s: %w", op, storage.ErrOrderRefunded)
	}

	if data[orderId].ClientID != clientId {
		return fmt.Errorf("%s: %w", op, storage.ErrClientNotOwner)
	}

	if data[orderId].IssueDate == "" {
		return fmt.Errorf("%s: %w", op, storage.ErrOrderNotIssued)
	}

	issueDate, err := time.Parse(layout, data[orderId].IssueDate)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	sub := time.Since(issueDate)
	dur := sub.Hours()
	var durDay int = int(dur / 24)

	if durDay >= 2 {
		return fmt.Errorf("%s: %w", op, storage.ErrMoreTwoDays)
	}

	o := data[orderId]
	o.RefundDate = time.Now().Format(layout)
	data[orderId] = o

	err = s.SaveData(&data)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil

}
