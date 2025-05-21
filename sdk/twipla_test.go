package sdk

import (
	"github.com/visitor-analytics/visa-3as-go-sdk/visa"
	"testing"
)

func TestEnv(t *testing.T) {
	if TwiplaProduction != visa.TwiplaEnvProduction {
		t.Error("env production mismatch")
	}
}
