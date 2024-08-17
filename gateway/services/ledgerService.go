package services

import (
	"example.com/m/gateway"
	"example.com/m/gateway/dto"
	"gorm.io/gorm"
)

type LedgerHandler struct {
	db *gorm.DB
}

func NewLedgerService(db *gorm.DB) *LedgerHandler {
	return &LedgerHandler{
		db: db,
	}
}

func (l *LedgerHandler) CreateLedgerRecord(dto dto.LedgerDto) (gateway.Ledger, error) {
	ledger := gateway.Ledger{
		TransactionId:          dto.TransactionId,
		AccountTo:              dto.AccountTo,
		AccountFrom:            dto.AccountFrom,
		InitialReceiverBalance: dto.InitialReceiverBalance,
		InitialSenderBalance:   dto.InitialSenderBalance,
		SenderBalance:          dto.SenderBalance,
		ReceiverBalance:        dto.ReceiverBalance,
		Amount:                 dto.Amount,
		RecordType:             dto.RecordType,
	}
	if err := l.db.Create(&ledger).Error; err != nil {
		return gateway.Ledger{}, err
	}
	return ledger, nil
}
