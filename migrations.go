package main

import (
	"example.com/m/gateway"
	"example.com/m/initializers"
	"example.com/m/wallet"
	"log"
)

func init() {
	initializers.StartLogger()
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(
		&wallet.User{},
		&wallet.Merchant{},
		&wallet.Profile{},
		&wallet.Account{},
		&gateway.Transaction{},
		&gateway.Ledger{},
		&gateway.MpesaTransaction{},
		&gateway.AirtelMoneyTransaction{},
	)
	if err != nil {
		log.Print(err)
	}
}
