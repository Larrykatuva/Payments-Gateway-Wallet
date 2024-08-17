package dto

import "example.com/m/gateway"

type RequestRrn struct{}

type RrnResponse struct {
	Rrn string
}

type LedgerDto struct {
	TransactionId          uint
	AccountFrom            *string
	AccountTo              *string
	InitialSenderBalance   *float64
	InitialReceiverBalance *float64
	SenderBalance          *float64
	ReceiverBalance        *float64
	Amount                 float64
	RecordType             gateway.RecordType
}
