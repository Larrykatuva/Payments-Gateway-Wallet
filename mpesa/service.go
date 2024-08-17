package mpesa

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"example.com/m/utils"
	"io"
	"net/http"
	"os"
	"time"
)

type Mpesa struct {
	username, password, baseUrl, shortCode, passKey string
}

func NewMpesaService(username, password, baseUrl, shortCode, passKey string) *Mpesa {
	return &Mpesa{
		username:  username,
		password:  password,
		baseUrl:   baseUrl,
		shortCode: shortCode,
		passKey:   passKey,
	}
}

func (m *Mpesa) handleResponse(format interface{}, response *http.Response) error {
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return errors.New(string(body))
	}
	err = json.Unmarshal(body, &format)
	if err != nil {
		return err
	}
	return nil
}

func (m *Mpesa) getToken(ctx context.Context) (TokenResult, error) {
	var token TokenResult
	url := m.baseUrl + "oauth/v1/generate?grant_type=client_credentials"
	response, err := utils.NewRequest(url, utils.POST).SetBasicAuth(utils.BasicAuth{
		Username: m.username,
		Password: m.password,
	}).Execute(ctx)
	if err != nil {
		return token, err
	}
	if err = m.handleResponse(token, response); err != nil {
		return token, err
	}
	return token, nil
}

func (m *Mpesa) generateTimestamp() string {
	currentTime := time.Now()
	return currentTime.Format("20060102150405")
}

func (m *Mpesa) generatePassword(timestamp string) string {
	dataToEncode := m.shortCode + m.passKey + timestamp
	encodedString := base64.StdEncoding.EncodeToString([]byte(dataToEncode))
	return encodedString
}

func (m *Mpesa) InitiateSTKRequest(ctx context.Context, amount, phoneNumber, description string) (StkpushResponse, error) {
	var stkResponse StkpushResponse
	url := m.baseUrl + "mpesa/stkpush/v1/processrequest"
	token, err := m.getToken(ctx)
	if err != nil {
		return stkResponse, err
	}
	timestamp := m.generateTimestamp()
	payload := StkInitiationPayload{
		BusinessShortCode: m.shortCode,
		Password:          m.generatePassword(timestamp),
		Timestamp:         timestamp,
		TransactionType:   PayBillOnline,
		Amount:            amount,
		PartyA:            phoneNumber,
		PartyB:            m.shortCode,
		PhoneNumber:       phoneNumber,
		CallBackURL:       os.Getenv("MPESA_CALLBACK"),
		AccountReference:  description,
		TransactionDesc:   description,
	}
	response, err := utils.NewRequest(url, utils.POST).SetBody(payload).SetBearerToken(token.AccessToken).Execute(ctx)
	if err != nil {
		return stkResponse, err
	}
	if err = m.handleResponse(stkResponse, response); err != nil {
		return stkResponse, err
	}
	return stkResponse, nil
}

func (m *Mpesa) QueryTransactionStatus(ctx context.Context, initiator, orderId, originatorId string) (StatusResult, error) {
	var status StatusResult
	url := m.baseUrl + "mpesa/transactionstatus/v1/query"
	token, err := m.getToken(ctx)
	if err != nil {
		return status, err
	}
	payload := StatusRequest{
		Initiator:                initiator,
		SecurityCredential:       m.generatePassword(m.generateTimestamp()),
		OriginatorConversationID: originatorId,
		CommandID:                QueryStatus,
		TransactionID:            orderId,
		PartyA:                   m.shortCode,
		IdentifierType:           "4",
		ResultURL:                os.Getenv("MPESA_QUERY_URL"),
		QueueTimeOutURL:          os.Getenv("MPESA_QUERY_URL"),
		Remarks:                  "OK",
		Occasion:                 "OK",
	}
	response, err := utils.NewRequest(url, utils.POST).SetBody(payload).SetBearerToken(token.AccessToken).Execute(ctx)
	if err != nil {
		return status, err
	}
	if err = m.handleResponse(status, response); err != nil {
		return status, err
	}
	return status, nil
}

func (m *Mpesa) InitiatePayout(ctx context.Context, initiator, orderId, amount, phoneNumber, description string) (C2BResponse, error) {
	var result C2BResponse
	url := m.baseUrl + "mpesa/b2c/v3/paymentrequest"
	token, err := m.getToken(ctx)
	if err != nil {
		return result, err
	}
	payload := C2BRequest{
		OriginatorConversationID: orderId,
		InitiatorName:            initiator,
		SecurityCredential:       m.generatePassword(m.generateTimestamp()),
		CommandID:                BusinessPayment,
		Amount:                   amount,
		PartyA:                   m.shortCode,
		PartyB:                   phoneNumber,
		Remarks:                  description,
		ResultURL:                os.Getenv("MPESA_QUERY_URL"),
		QueueTimeOutURL:          os.Getenv("MPESA_QUERY_URL"),
		Occasion:                 description,
	}
	response, err := utils.NewRequest(url, utils.POST).SetBody(payload).SetBearerToken(token.AccessToken).Execute(ctx)
	if err != nil {
		return result, err
	}
	if err = m.handleResponse(token, response); err != nil {
		return result, err
	}
	return result, nil
}

func (m *Mpesa) InitiateReversal(ctx context.Context, initiator, orderId, amount, phoneNumber, description string) (ReversalResponse, error) {
	var result ReversalResponse
	url := m.baseUrl + "mpesa/reversal/v1/request"
	token, err := m.getToken(ctx)
	if err != nil {
		return result, err
	}
	payload := ReversalRequest{
		Initiator:              initiator,
		SecurityCredential:     m.generatePassword(m.generateTimestamp()),
		CommandID:              Reversal,
		TransactionID:          orderId,
		Amount:                 amount,
		ReceiverParty:          phoneNumber,
		RecieverIdentifierType: "11",
		Remarks:                description,
		ResultURL:              os.Getenv("MPESA_QUERY_URL"),
		QueueTimeOutURL:        os.Getenv("MPESA_QUERY_URL"),
		Occasion:               description,
	}
	response, err := utils.NewRequest(url, utils.POST).SetBody(payload).SetBearerToken(token.AccessToken).Execute(ctx)
	if err != nil {
		return result, err
	}
	if err = m.handleResponse(result, response); err != nil {
		return result, err
	}
	return result, nil
}
