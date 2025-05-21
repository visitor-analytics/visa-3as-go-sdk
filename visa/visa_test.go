package visa

import (
	_ "embed"
	"fmt"
	"log"
	"testing"
)

//go:embed jwtRS256_intpc.key
var privateKey string

func ExampleTwiplaSDK_GenerateIframeURL() {
	// Create a new SDK instance.
	sdk, err := NewSDK(&TwiplaConfig{
		IntpID:      "2f8b7fd2-f958-4c10-b9d7-6aa0213ae299",
		PrivateKey:  privateKey,
		Environment: TwiplaEnvProduction,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Generate the dashboard URL for specific intpc ()
	url, err := sdk.GenerateIframeURL("first_test_01", "first_test_01_website_01")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(url)
}

func TestExample(t *testing.T) {
	ExampleTwiplaSDK_GenerateIframeURL()
}
