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

type TwiplaSSRApiClient struct {
	apiGateway string
	secret     string
}

func (t *TwiplaSSRApiClient) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", t.secret))

	return req, nil
}

func NewTwiplaSSRApiClient(apiGateway, secret string) *TwiplaSSRApiClient {
	return &TwiplaSSRApiClient{apiGateway: apiGateway, secret: secret}
}

type TwiplaApiClient struct {
	apiGateway string
	*AuthAPI
}

func (t *TwiplaApiClient) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
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

func NewTwiplaAPIClient(apiGateway string, jwt *AuthAPI) *TwiplaApiClient {
	return &TwiplaApiClient{apiGateway: apiGateway, AuthAPI: jwt}
}
