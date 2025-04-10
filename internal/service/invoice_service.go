package service

import (
	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
	"github.com/devfullcycle/imersao22/go-gateway/internal/dto"
)

type InvoiceService struct {
	invoiceRepository domain.InvoiceRepository
	accountService    AccountService
}

func NewInvoiceService(invoiceRepository domain.InvoiceRepository, accountService AccountService) *InvoiceService {
	return &InvoiceService{
		invoiceRepository: invoiceRepository,
		accountService:    accountService,
	}
}

func (s *InvoiceService) Create(input *dto.CreateInvoceInput) (*dto.InvoiceOutPut, error) {
	accountOutPut, err := s.accountService.FindByAPIKey(input.APIKey)
	if err != nil {
		return nil, err
	}

	invoice, err := dto.ToInvoice(input, accountOutPut.ID)
	if err != nil {
		return nil, err
	}

	if err = invoice.Process(); err != nil {
		return nil, err
	}

	if invoice.Status == domain.StatusApproved {
		_, err = s.accountService.UpdateBalance(input.APIKey, invoice.Amount)
		if err != nil {
			return nil, err
		}
	}

	if err := s.invoiceRepository.Save(invoice); err != nil {
		return nil, err
	}

	return dto.FromInvoice(invoice), nil

}

func (s *InvoiceService) GetByID(id, apikey string) (*dto.InvoiceOutPut, error) {
	invoice, err := s.invoiceRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	accountOutPut, err := s.accountService.FindByAPIKey(apikey)
	if err != nil {
		return nil, err
	}

	if invoice.AccountID != accountOutPut.ID {
		return nil, domain.ErrUnauthorizedAccess
	}

	return dto.FromInvoice(invoice), nil
}

func (s *InvoiceService) ListByAccount(accountID string) ([]*dto.InvoiceOutPut, error) {
	invoices, err := s.invoiceRepository.FindByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	output := make([]*dto.InvoiceOutPut, len(invoices))
	for i, invoice := range invoices {
		output[i] = dto.FromInvoice(invoice)
	}

	return output, nil

}

func (s *InvoiceService) ListByAccountAPIKey(apikey string) ([]*dto.InvoiceOutPut, error) {
	accountOutPut, err := s.accountService.FindByAPIKey(apikey)
	if err != nil {
		return nil, domain.ErrAccountNotFound
	}

	return s.ListByAccount(accountOutPut.ID)
}
