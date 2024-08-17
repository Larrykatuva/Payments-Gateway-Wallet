package gateway

import (
	"example.com/m/wallet"
	"gorm.io/gorm"
)

type TransactionStatus string

type TransactionProvider string

type ServiceType string

type RecordType string

const (
	Pending         TransactionStatus   = "Pending"
	Successful      TransactionStatus   = "Successful"
	Failed          TransactionStatus   = "Failed"
	Reversed        TransactionStatus   = "Reversed"
	Kimipay         TransactionProvider = "Kimipay"
	Mpesa           TransactionProvider = "Mpesa"
	AirtelMoney     TransactionProvider = "AirtelMoney"
	Topup           ServiceType         = "Topup"
	MerchantPayment ServiceType         = "MerchantPayment"
	Payout          ServiceType         = "Payout"
	InitialRecord   RecordType          = "InitialRecord"
	TransactionFee  RecordType          = "TransactionFee"
)

type Transaction struct {
	gorm.Model
	MerchantId             uint
	Merchant               wallet.Merchant
	ProviderRef            *string
	Provider               TransactionProvider
	ServiceType            ServiceType
	Rrn                    string `gorm:"unique"`
	OrderId                string
	Description            string
	AccountFrom            *string
	AccountTo              *string
	InitialSenderBalance   *string
	InitialReceiverBalance *string
	SenderBalance          *float64
	ReceiverBalance        *float64
	Status                 TransactionStatus
	Completed              bool `gorm:"default:false"`
	Amount                 float64
	Currency               wallet.Currency
	Metadata               *string
}

type Ledger struct {
	gorm.Model
	TransactionId          uint
	Transaction            Transaction
	AccountFrom            *string
	AccountTo              *string
	InitialSenderBalance   *float64
	InitialReceiverBalance *float64
	SenderBalance          *float64
	ReceiverBalance        *float64
	Amount                 float64
	RecordType             RecordType
}

type MpesaTransaction struct {
	gorm.Model
	TransactionId     uint
	Transaction       Transaction
	MerchantRequestId string
	CheckoutRequestId string
	PhoneNumber       string
	Names             *string
	Ref               *string
	ResultCode        *int
	ResultDescription *string
	BusinessShortCode string
	Status            TransactionStatus
}

type AirtelMoneyTransaction struct {
	gorm.Model
	TransactionId     uint
	Transaction       Transaction
	PartnerCode       string
	ResponseCode      string
	ResponseStatus    string
	ResultDescription string
	MoneyId           string
	Status            TransactionStatus
}
