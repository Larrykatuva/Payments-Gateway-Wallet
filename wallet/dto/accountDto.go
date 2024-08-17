package wallet

import "example.com/m/wallet"

type ActionType string

const (
	Debit  ActionType = "Debit"
	Credit ActionType = "Credit"
	Hold   ActionType = "Hold"
)

type AccountDto struct {
	Name, Description string
	Type              wallet.AccountType
}

type CreateAccountRequest struct {
	AccountDto
	User     wallet.User
	Merchant wallet.Merchant
}

type CreateAccountResponse struct {
	Account wallet.Account
	Err     error
}

type ActionRequest struct {
	AccountNumber string
	Amount        float64
	Action        ActionType
}

type ActionResponse struct {
	Status bool
	Err    error
}
