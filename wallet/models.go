package wallet

import "gorm.io/gorm"

type UserType string
type AccountType string
type Gender string
type Currency string

const (
	Individual UserType    = "Individual"
	Corporate  UserType    = "Corporate"
	Personal   AccountType = "Personal"
	Business   AccountType = "Individual"
	Male       Gender      = "Male"
	Female     Gender      = "Female"
	Other      Gender      = "Other"
	Kes        Currency    = "Kes"
)

type User struct {
	gorm.Model
	IdentityType   string
	Type           UserType
	Sub            string
	IdentityNumber string `gorm:"uniqueIndex"`
	PhoneNumber    string `gorm:"uniqueIndex;size:12"`
	Email          string `gorm:"uniqueIndex"`
	Address        string
}

type Profile struct {
	gorm.Model
	FirstName   string
	LastName    string
	Gender      Gender
	DateOfBirth string
	County      string
	SubCounty   string
	UserID      uint
	User        User
}

type Merchant struct {
	gorm.Model
	BusinessName       string
	Email              string
	RegistrationNumber string
	Description        string
	Sub                string
}

type Account struct {
	gorm.Model
	Name           string
	Description    string
	AccountNumber  string
	Currency       Currency
	Type           AccountType
	Balance        float64
	RunningBalance float64
	Active         bool `gorm:"default:true"`
	UserID         uint
	User           User
	MerchantID     uint
	Merchant       Merchant
}
