package domain

import "time"

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

type Invoice struct {
	ID             string
	AccountID      string
	Status         Status
	Description    string
	PaymentsType   string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
type CreditCard struct {
	Number         string
	CVV            string
	ExpiryMonth    int
	ExpiryYear     int
	CardholderName string
}
