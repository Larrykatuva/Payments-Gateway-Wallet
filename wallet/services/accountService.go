package services

import (
	"errors"
	"example.com/m/wallet"
	dto "example.com/m/wallet/dto"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

const (
	InsufficientBalance = "insufficient account balance"
	UnknownAction       = "unknown action request"
)

type AccountHandler struct {
	db *gorm.DB
}

func NewAccountService(db *gorm.DB) *AccountHandler {
	return &AccountHandler{
		db: db,
	}
}

func (a *AccountHandler) generateAccountNumber() (string, error) {
	var account wallet.Account
	if err := a.db.Last(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "1000000", nil
		} else {
			return "", err
		}
	}
	newAccountNumber, err := strconv.ParseUint(account.AccountNumber, 10, 64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", newAccountNumber), nil
}

func (a *AccountHandler) CreateAccount(dto dto.AccountDto, user wallet.User, merchant wallet.Merchant) (wallet.Account, error) {
	accountNumber, err := a.generateAccountNumber()
	if err != nil {
		return wallet.Account{}, err
	}
	account := wallet.Account{
		Name:           dto.Name,
		Description:    dto.Description,
		AccountNumber:  accountNumber,
		Currency:       wallet.Kes,
		Type:           dto.Type,
		Balance:        0.00,
		RunningBalance: 0.00,
		Active:         true,
		UserID:         user.ID,
		MerchantID:     merchant.ID,
	}
	if err = a.db.Create(&account).Error; err != nil {
		return wallet.Account{}, err
	}
	return account, nil
}

func (a *AccountHandler) FilterAccounts(filters wallet.Account) ([]wallet.Account, error) {
	var accounts []wallet.Account
	result := a.db.Where(filters).Find(&accounts)
	if result.RowsAffected == 0 {
		return accounts, nil
	}
	if result.Error != nil {
		return accounts, result.Error
	}
	return accounts, nil
}

func (a *AccountHandler) FilterAccount(filters wallet.Account) (*wallet.Account, error) {
	var account wallet.Account
	result := a.db.Where(filters).First(&account)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

func (a *AccountHandler) HoldInTransactionBalance(dto dto.ActionRequest) error {
	account, err := a.FilterAccount(wallet.Account{AccountNumber: dto.AccountNumber})
	if err != nil {
		return err
	}
	if account.Balance < dto.Amount {
		return errors.New(InsufficientBalance)
	}
	account.Balance = account.Balance - dto.Amount
	account.RunningBalance = account.RunningBalance + dto.Amount
	a.db.Save(&account)
	return nil
}

func (a *AccountHandler) DebitAccount(dto dto.ActionRequest) error {
	account, err := a.FilterAccount(wallet.Account{AccountNumber: dto.AccountNumber})
	if err != nil {
		return err
	}
	account.RunningBalance = account.RunningBalance - dto.Amount
	a.db.Save(&account)
	return nil
}

func (a *AccountHandler) CreditAccount(dto dto.ActionRequest) error {
	account, err := a.FilterAccount(wallet.Account{AccountNumber: dto.AccountNumber})
	if err != nil {
		return err
	}
	account.Balance = account.Balance + dto.Amount
	a.db.Save(&account)
	return nil
}

func (a *AccountHandler) Handler(payload dto.ActionRequest) error {
	switch payload.Action {
	case dto.Hold:
		return a.HoldInTransactionBalance(payload)
	case dto.Debit:
		return a.DebitAccount(payload)
	case dto.Credit:
		return a.CreditAccount(payload)
	default:
		return errors.New(UnknownAction)
	}
}
