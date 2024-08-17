package dto

import (
	"example.com/m/gateway"
	"example.com/m/mpesa"
	"example.com/m/wallet"
)

type CreateTransactionDto struct {
	MerchantId  uint
	Provider    gateway.TransactionProvider
	ServiceType gateway.ServiceType
	Rrn         string
	OrderId     string
	AccountFrom *string
	AccountTo   *string
	Amount      float64
	Currency    wallet.Currency
}

type UpdateTransactionDto struct {
	ProviderRef            *string
	SenderBalance          *float64
	ReceiverBalance        *float64
	Status                 *gateway.TransactionStatus
	Completed              *bool
	InitialSenderBalance   *string
	InitialReceiverBalance *string
	Description            *string
}

type MpesaConfig struct {
	ShortCode, PassKey, Username, Password string
}

type MpesaTopupRequest struct {
	TransactionId uint
	AccountTo     string
	PhoneNumber   string
	Amount        float64
	MpesaCallback mpesa.StkCallback
}

type TransactionFee struct {
	Amount    float64
	AccountTo string
}

type WalletTransferRequest struct {
	TransactionId uint
	AccountFrom   string
	AccountTo     string
	Amount        float64
	Fee           *TransactionFee
}
