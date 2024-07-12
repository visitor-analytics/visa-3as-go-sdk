package sdk

import (
	"fmt"
	"io"
	"net/http"
)

type PagArgs struct {
	Page     int
	PageSize int
}

type TwiplaAPIClient struct {
	apiGateway string
	*AuthAPI
}

func (t *TwiplaAPIClient) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	accessToken, err := t.AuthAPI.NewINTPToken()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	return req, nil
}

func NewTwiplaAPIClient(apiGateway string, jwt *AuthAPI) *TwiplaAPIClient {
	return &TwiplaAPIClient{apiGateway: apiGateway, AuthAPI: jwt}
}
