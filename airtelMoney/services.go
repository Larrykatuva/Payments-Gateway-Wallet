package airtelMoney

import (
	"context"
	"encoding/json"
	"errors"
	"example.com/m/utils"
	"io"
	"net/http"
)

type AirtelMoney struct {
	clientId, clientSecret, baseUrl string
}

func NewAirtelMoneyService(clientId, clientSecret, baseUrl string) *AirtelMoney {
	return &AirtelMoney{
		clientId:     clientId,
		clientSecret: clientSecret,
		baseUrl:      baseUrl,
	}
}

func (a *AirtelMoney) handleResponse(format interface{}, response *http.Response) error {
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

func (a *AirtelMoney) getToken(ctx context.Context) (TokenResult, error) {
	var result TokenResult
	url := a.baseUrl + "auth/oauth2/token"
	response, err := utils.NewRequest(url, utils.POST).SetBody(TokenRequest{
		ClientID:     a.clientId,
		ClientSecret: a.clientSecret,
		GrantType:    "client_credentials",
	}).Execute(ctx)
	if err != nil {
		return result, err
	}
	if err = a.handleResponse(result, response); err != nil {
		return result, err
	}
	return result, nil
}
