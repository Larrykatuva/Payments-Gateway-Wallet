package actors

import (
	"example.com/m/gateway"
	gatewayDto "example.com/m/gateway/dto"
	gatewayServices "example.com/m/gateway/services"
	"example.com/m/initializers"
	walletModels "example.com/m/wallet"
	walletDto "example.com/m/wallet/dto"
	walletServices "example.com/m/wallet/services"
	"fmt"
	"github.com/anthdm/hollywood/actor"
)

type AccountManager struct{}

var Engine *actor.Engine

var accountStore = make(map[string]*actor.PID)

func NewAccountManager() actor.Receiver {
	return &AccountManager{}
}

func GetAccountPid(account string) *actor.PID {
	if accountStore[account] != nil {
		return accountStore[account]
	}
	accountManager := NewAccountManager
	pid := Engine.Spawn(accountManager, account)
	accountStore[account] = pid
	return pid
}

func (a *AccountManager) Receive(context *actor.Context) {
	switch msg := context.Message().(type) {
	case walletDto.CreateAccountRequest:
		accountHandler := walletServices.NewAccountService(initializers.DB)
		account, err := accountHandler.CreateAccount(walletDto.AccountDto{
			Name:        msg.Name,
			Description: msg.Description,
			Type:        msg.Type,
		}, msg.User, msg.Merchant)
		context.Respond(walletDto.CreateAccountResponse{
			Account: account,
			Err:     err,
		})
	case walletDto.ActionRequest:
		accountHandler := walletServices.NewAccountService(initializers.DB)
		err := accountHandler.Handler(msg)
		response := walletDto.ActionResponse{
			Status: true,
			Err:    err,
		}
		if err != nil {
			response.Status = false
		}
		context.Respond(response)
	case gatewayDto.MpesaTopupRequest:
		accountHandler := walletServices.NewAccountService(initializers.DB)
		account, _ := accountHandler.FilterAccount(walletModels.Account{AccountNumber: msg.AccountTo})
		receiverBalance := account.Balance + msg.Amount
		err := accountHandler.CreditAccount(walletDto.ActionRequest{
			AccountNumber: account.AccountNumber,
			Amount:        msg.Amount,
			Action:        walletDto.Credit,
		})
		if err != nil {
			fmt.Print(err)
		}
		_, err = gatewayServices.NewLedgerService(initializers.DB).CreateLedgerRecord(gatewayDto.LedgerDto{
			TransactionId:          msg.TransactionId,
			AccountFrom:            &msg.PhoneNumber,
			AccountTo:              &msg.AccountTo,
			InitialReceiverBalance: &account.Balance,
			ReceiverBalance:        &receiverBalance,
			Amount:                 msg.Amount,
			RecordType:             gateway.InitialRecord,
		})
		if err != nil {
			fmt.Print(err)
		}
	case gatewayDto.WalletTransferRequest:
		accountHandler := walletServices.NewAccountService(initializers.DB)
		accountFrom, _ := accountHandler.FilterAccount(walletModels.Account{AccountNumber: msg.AccountFrom})
		accountTo, _ := accountHandler.FilterAccount(walletModels.Account{AccountNumber: msg.AccountTo})
		err := accountHandler.DebitAccount(walletDto.ActionRequest{
			AccountNumber: accountFrom.AccountNumber,
			Amount:        msg.Amount,
			Action:        walletDto.Debit,
		})
		if err != nil {
			fmt.Print(err)
		}
		err = accountHandler.CreditAccount(walletDto.ActionRequest{
			AccountNumber: accountTo.AccountNumber,
			Amount:        msg.Amount,
			Action:        walletDto.Credit,
		})
		if err != nil {
			fmt.Print(err)
		}
		InitialSenderBalance := accountFrom.Balance + msg.Amount
		receiverBalance := accountTo.Balance + msg.Amount
		_, err = gatewayServices.NewLedgerService(initializers.DB).CreateLedgerRecord(gatewayDto.LedgerDto{
			TransactionId:          msg.TransactionId,
			AccountFrom:            &msg.AccountFrom,
			AccountTo:              &msg.AccountTo,
			InitialSenderBalance:   &InitialSenderBalance,
			InitialReceiverBalance: &accountTo.Balance,
			SenderBalance:          &accountFrom.Balance,
			ReceiverBalance:        &receiverBalance,
			Amount:                 msg.Amount,
			RecordType:             gateway.InitialRecord,
		})
	}
}
