package dto

import (
	"time"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
)

const (
	StatusPending  = string(domain.StatusPending)
	StatusApproved = string(domain.StatusApproved)
	StatusRejected = string(domain.StatusRejected)
)

type CreateInvoceInput struct {
	APIKey         string
	Amount         float64 `json:"amount"`
	Description    string  `json:"description"`
	PaymentsType   string  `json:"payments_type"`
	CardNumber     string  `json:"card_number"`
	CVV            string  `json:"cvv"`
	ExpiryMonth    int     `json:"expiry_month"`
	ExpiryYear     int     `json:"expiry_year"`
	CardholderName string  `json:"cardholder_name"`
}

type InvoiceOutPut struct {
	ID             string    `json:"id"`
	AccountID      string    `json:"account_id"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"`
	Description    string    `json:"description"`
	PaymentsType   string    `json:"payments_type"`
	CardLastDigits string    `json:"card_last_digits"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ToInvoice(input *CreateInvoceInput, accountID string) (*domain.Invoice, error) {
	card := domain.CreditCard{
		Number:         input.CardNumber,
		CVV:            input.CVV,
		ExpiryMonth:    input.ExpiryMonth,
		ExpiryYear:     input.ExpiryYear,
		CardholderName: input.CardholderName,
	}

	return domain.NewInvoice(
		accountID,
		input.Amount,
		input.Description,
		input.PaymentsType,
		card,
	)
}

func FromInvoice(invoice *domain.Invoice) *InvoiceOutPut {
	return &InvoiceOutPut{
		ID:             invoice.ID,
		AccountID:      invoice.AccountID,
		Amount:         invoice.Amount,
		Status:         string(invoice.Status),
		Description:    invoice.Description,
		PaymentsType:   invoice.PaymentsType,
		CardLastDigits: invoice.CardLastDigits,
		CreatedAt:      invoice.CreatedAt,
		UpdatedAt:      invoice.UpdatedAt,
	}
}
