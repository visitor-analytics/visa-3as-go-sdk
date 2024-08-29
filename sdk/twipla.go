package sdk

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type TwiplaEnv string
type TwiplaCurrency string
type TwiplaPeriod string

const TwiplaDevelop TwiplaEnv = "dev"
const TwiplaStage TwiplaEnv = "stage"
const TwiplaProduction TwiplaEnv = "production"

type Intp struct {
	ID   string
	PKey string
}

type SSR struct {
	Secret string
}

type TwiplaArgs struct {
	Intp Intp
	Env  TwiplaEnv
	Ssr  *SSR
}

type Twipla struct {
	Auth         *AuthAPI
	Intpc        *TwiplaIntpcAPI
	Website      *TwiplaWebsiteAPI
	Package      *TwiplaPackageAPI
	Subscription *TwiplaSubscriptionAPI
}

func NewTwipla(args TwiplaArgs) (*Twipla, error) {
	var apiGateway string
	var apiGatewaySSR string

	switch args.Env {
	case TwiplaDevelop:
		{
			apiGateway = "https://api-gateway.va-endpoint.com"
			apiGatewaySSR = "https://dev-api.va-endpoint.com"
		}

	case TwiplaStage:
		{
			apiGateway = "https://stage-api-gateway.va-endpoint.com"
			apiGatewaySSR = "https://stage-api.va-endpoint.com"
		}

	case TwiplaProduction:
		{
			apiGateway = "https://api-gateway.visitor-analytics.io"
			apiGatewaySSR = "https://lb-api.visitor-analytics.io"
		}

	default:
		return nil, fmt.Errorf("unsupported env: %s", args.Env)
	}

	var twiplaSSRWebsiteAPI *TwiplaSSRWebsiteAPI
	if args.Ssr != nil {
		twiplaSSRWebsiteAPI = NewTwiplaSSRWebsiteAPI(NewTwiplaSSRApiClient(apiGatewaySSR, args.Ssr.Secret))
	}

	pkey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(args.Intp.PKey))
	if err != nil {
		return nil, err
	}

	authAPI := NewAuthAPI(args.Intp.ID, &RS256{PrivateKey: pkey})
	twiplaAPIClient := NewTwiplaAPIClient(apiGateway, authAPI)
	websiteSubscriptionAPI := NewTwiplaSubscriptionAPI(twiplaAPIClient)
	intpPackageAPI := NewTwiplaIntpPackageAPI(twiplaAPIClient)
	websiteAPI := NewTwiplaWebsiteAPI(twiplaAPIClient, twiplaSSRWebsiteAPI)
	intpcAPI := NewTwiplaIntpcAPI(twiplaAPIClient, websiteAPI, twiplaSSRWebsiteAPI)

	return &Twipla{
		Auth:         authAPI,
		Intpc:        intpcAPI,
		Website:      websiteAPI,
		Package:      intpPackageAPI,
		Subscription: websiteSubscriptionAPI,
	}, nil
}
