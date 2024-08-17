package wallet

import (
	"example.com/m/wallet"
)

type UserDto struct {
	IdentityType   string
	Type           wallet.UserType
	Sub            string
	IdentityNumber string
	PhoneNumber    string
	Email          string
	Address        string
}

type ProfileDto struct {
	FirstName   string
	LastName    string
	Gender      wallet.Gender
	DateOfBirth string
	County      string
	SubCounty   string
}

type MerchantDto struct {
	BusinessName       string
	Email              string
	RegistrationNumber string
	Description        string
	Sub                string
}
