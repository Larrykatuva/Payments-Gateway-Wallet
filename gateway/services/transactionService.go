package services

import (
	"example.com/m/gateway"
	"example.com/m/gateway/dto"
	"gorm.io/gorm"
)

const InitialDescription = "Transaction processing"

type TransactionHandler struct {
	db *gorm.DB
}

func NewTransactionService(db *gorm.DB) *TransactionHandler {
	return &TransactionHandler{
		db: db,
	}
}

func (t *TransactionHandler) CreateTransaction(dto dto.CreateTransactionDto) (gateway.Transaction, error) {
	transaction := gateway.Transaction{
		MerchantId:  dto.MerchantId,
		Provider:    dto.Provider,
		ServiceType: dto.ServiceType,
		Description: InitialDescription,
		Rrn:         dto.Rrn,
		OrderId:     dto.OrderId,
		AccountFrom: dto.AccountFrom,
		AccountTo:   dto.AccountTo,
		Status:      gateway.Pending,
		Amount:      dto.Amount,
		Currency:    dto.Currency,
	}
	if err := t.db.Create(&transaction).Error; err != nil {
		return gateway.Transaction{}, err
	}
	return transaction, nil
}

func (t *TransactionHandler) FilterTransaction(filters gateway.Transaction) (*gateway.Transaction, error) {
	var transaction gateway.Transaction
	result := t.db.Where(filters).First(&transaction)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &transaction, nil
}

func (t *TransactionHandler) UpdateTransaction(filters gateway.Transaction, payload dto.UpdateTransactionDto) (*gateway.Transaction, error) {
	transaction, err := t.FilterTransaction(filters)
	if err != nil {
		return nil, err
	}
	if payload.Status != nil {
		transaction.Status = *payload.Status
	}
	if payload.ProviderRef != nil {
		transaction.ProviderRef = payload.ProviderRef
	}
	if payload.Completed != nil {
		transaction.Completed = *payload.Completed
	}
	if payload.ReceiverBalance != nil {
		transaction.ReceiverBalance = payload.ReceiverBalance
	}
	if payload.SenderBalance != nil {
		transaction.SenderBalance = payload.SenderBalance
	}
	if payload.InitialSenderBalance != nil {
		transaction.InitialSenderBalance = payload.InitialSenderBalance
	}
	if payload.InitialReceiverBalance != nil {
		transaction.InitialReceiverBalance = payload.InitialReceiverBalance
	}
	if payload.Description != nil {
		transaction.Description = *payload.Description
	}
	t.db.Save(&transaction)
	return t.FilterTransaction(filters)
}
