package sdk

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func NewAuthAPI(intpID string, signer TokenSigner) *AuthAPI {
	return &AuthAPI{
		intpID,
		signer,
	}
}

type AuthAPI struct {
	IntpID string
	Signer TokenSigner
}

func (t *AuthAPI) NewINTPToken() (string, error) {
	return t.Signer.Sign(t.IntpID, jwt.MapClaims{
		"iss":     "twipla-3as-go-sdk",
		"roles":   []string{"intp"},
		"intp_id": t.IntpID,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 4).Unix(),
	})
}

func (t *AuthAPI) NewINTPcToken(intpcID string) (string, error) {
	return t.Signer.Sign(t.IntpID, jwt.MapClaims{
		"iss":      "twipla-3as-go-sdk",
		"roles":    []string{"intpc"},
		"intp_id":  t.IntpID,
		"intpc_id": intpcID,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 4).Unix(),
	})
}
