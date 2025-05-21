package visa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"net/url"
)

type TwiplaEnv string

const (
	TwiplaEnvDevelop    TwiplaEnv = "dev"
	TwiplaEnvStage      TwiplaEnv = "stage"
	TwiplaEnvProduction TwiplaEnv = "production"
)

type TwiplaConfig struct {
	// IntpID is the ID issued by TWIPLA for an INTP, in exchange for `jwtRS256.key.pub`
	IntpID string
	// PrivateKey contains the plaintext contents of the PEM file containing the INTP's private key
	PrivateKey string

	// Environment sets which TWIPLA deployment to use.
	Environment TwiplaEnv
}

type TwiplaSDK struct {
	signer  *tokenSigner
	client  *http.Client
	env     TwiplaEnv
	apiBase *url.URL
}

func (sdk *TwiplaSDK) apiCall(ctx context.Context, method string, path string, body any) (*http.Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequestWithContext(ctx, method, sdk.apiBase.JoinPath(path).String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	token, err := sdk.signer.IntpToken()
	if err != nil {
		return nil, fmt.Errorf("can't sign bearer intp token: %w", err)
	}
	r.Header.Set("Authorization", "Bearer "+token)
	if !(method == http.MethodGet || method == http.MethodDelete) {
		r.Header.Set("Content-Type", "application/json")
	}

	return http.DefaultClient.Do(r)
}

func NewSDK(config *TwiplaConfig) (*TwiplaSDK, error) {

	pkey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(config.PrivateKey))
	if err != nil {
		return nil, err
	}

	signer := &tokenSigner{
		privateKey: pkey,
		intpID:     config.IntpID,
	}

	var apiPrefix string
	switch config.Environment {
	case TwiplaEnvDevelop:
		apiPrefix = "https://api-gateway.va-endpoint.com"
	case TwiplaEnvStage:
		apiPrefix = "https://stage-api-gateway.va-endpoint.com"
	default:
		apiPrefix = "https://api-gateway.visitor-analytics.io"
		config.Environment = TwiplaEnvProduction
	}

	apiURL, err := url.Parse(apiPrefix)

	return &TwiplaSDK{
		signer:  signer,
		env:     config.Environment,
		apiBase: apiURL,
	}, nil
}
