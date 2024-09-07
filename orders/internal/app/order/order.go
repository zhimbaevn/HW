package order

import "fmt"

type Order struct {
	OrderId       int    `json:"order_id"`
	ClientID      int    `json:"client_id"`
	StoragePeriod string `json:"storage_period"`
	IssueDate     string `json:"issue_date"`
	RefundDate    string `json:"refund_date"`
	RecTime       string `json:"recording_time"`
}

func (o Order) String() string {
	return fmt.Sprintf("order id: %d, client id: %d, storage period: %s, issue date: %s, refund date: %s", o.OrderId, o.ClientID, o.StoragePeriod, o.IssueDate, o.RefundDate)
}
