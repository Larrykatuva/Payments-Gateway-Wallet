package airtelMoney

import "time"

type UssdPushResponse struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

type Subscriber struct {
	Country  string `json:"country"`
	Currency string `json:"currency"`
	Msisdn   string `json:"msisdn"`
}

type TransactionResponse struct {
	Amount   int    `json:"amount"`
	Country  string `json:"country"`
	Currency string `json:"currency"`
	Id       string `json:"id"`
}

type UssdPushRequest struct {
	Reference   string              `json:"reference"`
	Subscriber  Subscriber          `json:"subscriber"`
	Transaction TransactionResponse `json:"transaction"`
}

type UssdPushBill struct {
	PhoneNumber string `json:"phoneNumber"`
	Amount      int    `json:"amount"`
	Reference   string `json:"reference"`
	CallbackURL string `json:"callbackurl"`
	From        string `json:"from"`
}

type TokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type,omitempty"`
}

type TokenResult struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type CallbackPayload struct {
	PhoneNumber     string    `json:"phoneNumber"`
	TransactionRef  string    `json:"transactionRef"`
	Amount          int       `json:"amount"`
	Status          string    `json:"status"`
	TransactionTime time.Time `json:"transactionTime"`
	FailedReason    string    `json:"failedReason,omitempty"`
	Message         string    `json:"message,omitempty"`
}
