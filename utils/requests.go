package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	POST             = "POST"
	GET              = "GET"
	PATCH            = "PATCH"
	JSON             = "application/json"
	FORM_URL_ENCODED = "application/x-www-form-urlencoded"
)

type BasicAuth struct {
	Username, Password string
}

type RequestAdapter struct {
	Url, Method, ContentType, BearerToken string
	Data                                  interface{}
	FormData                              url.Values
	BasicAuth                             BasicAuth
}

func NewRequest(url string, method string) *RequestAdapter {
	return &RequestAdapter{
		Url:         url,
		Method:      method,
		ContentType: JSON,
	}
}

func (rq *RequestAdapter) SetBasicAuth(auth BasicAuth) *RequestAdapter {
	rq.BasicAuth = auth
	return rq
}

func (rq *RequestAdapter) SetContentType(contentType string) *RequestAdapter {
	rq.ContentType = contentType
	return rq
}

func (rq *RequestAdapter) SetFormData(formData url.Values) *RequestAdapter {
	rq.FormData = formData
	return rq
}

func (rq *RequestAdapter) SetBody(body interface{}) *RequestAdapter {
	rq.Data = body
	return rq
}

func (rq *RequestAdapter) SetBearerToken(token string) *RequestAdapter {
	rq.BearerToken = token
	return rq
}

func generateRequest(ctx context.Context, rq *RequestAdapter) (*http.Request, error) {
	var request *http.Request
	var err error
	if rq.Method == GET {
		request, err = http.NewRequestWithContext(ctx, rq.Method, rq.Url, nil)
	} else {
		if rq.ContentType == JSON {
			var jsonData []byte
			jsonData, err = json.Marshal(rq.Data)
			if err != nil {
				return nil, err
			}
			request, err = http.NewRequestWithContext(ctx, rq.Method, rq.Url, bytes.NewBuffer(jsonData))
		} else {
			request, err = http.NewRequestWithContext(ctx, rq.Method, rq.Url, bytes.NewBufferString(rq.FormData.Encode()))
		}
	}
	return request, nil
}

func (rq *RequestAdapter) Execute(ctx context.Context) (*http.Response, error) {
	request, err := generateRequest(ctx, rq)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", rq.ContentType)
	if len(rq.BearerToken) != 0 {
		request.Header.Set("Authorization", "Bearer "+rq.BearerToken)
	}
	if rq.BasicAuth.Username != "" {
		request.SetBasicAuth(rq.BasicAuth.Username, rq.BasicAuth.Password)
	}
	client := &http.Client{}
	return client.Do(request)
}
