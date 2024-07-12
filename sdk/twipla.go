package sdk

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type TwiplaEnv string
type TwiplaCurrency string
type TwiplaPeriod string

const Monthly TwiplaPeriod = "monthly"
const Yearly TwiplaPeriod = "yearly"

const EUR TwiplaCurrency = "EUR"
const USD TwiplaCurrency = "USD"
const RON TwiplaCurrency = "RON"

const TwiplaDevelop TwiplaEnv = "dev"
const TwiplaStage TwiplaEnv = "stage"
const TwiplaProduction TwiplaEnv = "production"

type Intp struct {
	ID   string
	PKey string
}

type TwiplaArgs struct {
	Intp Intp
	Env  TwiplaEnv
}

type Twipla struct {
	Auth         *AuthAPI
	Package      *TwiplaPackageAPI
	Website      *TwiplaWebsiteAPI
	Customer     *TwiplaCustomerAPI
	Subscription *TwiplaSubscriptionAPI
}

func NewTwipla(args TwiplaArgs) (*Twipla, error) {
	var apiGateway string

	switch args.Env {
	case TwiplaDevelop:
		apiGateway = "https://api-gateway.va-endpoint.com"
	case TwiplaStage:
		apiGateway = "https://stage-api-gateway.va-endpoint.com"
	case TwiplaProduction:
		apiGateway = "https://api-gateway.visitor-analytics.io"
	default:
		return nil, fmt.Errorf("unsupported env: %s", args.Env)
	}

	pkey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(args.Intp.PKey))
	if err != nil {
		return nil, err
	}

	authAPI := NewAuthAPI(args.Intp.ID, &RS256{PrivateKey: pkey})

	twiplaAPIClient := NewTwiplaAPIClient(apiGateway, authAPI)

	return &Twipla{
		Auth:         authAPI,
		Package:      NewTwiplaPackageAPI(twiplaAPIClient),
		Website:      NewTwiplaWebsiteAPI(twiplaAPIClient),
		Customer:     NewTwiplaCustomerAPI(twiplaAPIClient),
		Subscription: NewTwiplaSubscriptionAPI(twiplaAPIClient),
	}, nil
}
