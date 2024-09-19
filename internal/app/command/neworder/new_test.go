package neworder

import (
	"errors"
	"orders/internal/storage"
	"orders/internal/storage/jsondb"
	"testing"
)

const storage_path = "C:/GoLess/myHomeWork/orders/storage/storage.json"

func TestNew_New(t *testing.T) {
	s, err := jsondb.GetStorage(storage_path)
	if err != nil {
		t.Fatalf("TestNew_New: %s", err)
	}
	ErrStoragePeriodExpired := errors.New("the order storage period has already expired")

	cases := []struct {
		orderId       int
		clientId      int
		storagePeriod string
		retErr        error
	}{
		{
			orderId:       1337,
			clientId:      123,
			storagePeriod: "2024-09-09",
			retErr:        nil,
		},
		{
			orderId:       1337,
			clientId:      123,
			storagePeriod: "2024-09-09",
			retErr:        storage.ErrOrderExist,
		},
		{
			orderId:       12345,
			clientId:      123,
			storagePeriod: "2024-09-01",
			retErr:        ErrStoragePeriodExpired,
		},
		{
			orderId:       54312,
			clientId:      123,
			storagePeriod: "2024-20034-01",
			retErr:        errors.New("month out of range"),
		},
	}

	// TODO CASES:
	// 1. Рабочий кейс с добавлением | return nil
	// 2. Срок хранения истёк | return ErrStoragePeriodExpired
	// 3. Неправильный формат даты
	// 4. Заказ существует | return ErrOrderExist

	for _, tc := range cases {

		result := New(s, tc.orderId, tc.clientId, tc.storagePeriod)

		t.Logf("Calling New(%s, %d, %d, %s), result: %s\n", s, tc.orderId, tc.clientId, tc.storagePeriod, result)
		if tc.retErr != nil {
			if errors.Is(err, tc.retErr) {
				t.Errorf("Incorrect result. Expect: %s, got: %s", tc.retErr, result)
			}
		} else {
			if err != tc.retErr {
				t.Errorf("Incorrect result. Expect: %s, got: %s", tc.retErr, result)
			}
		}

	}

}
