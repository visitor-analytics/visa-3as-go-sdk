package sdk

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TwiplaJWTIssuer struct {
	intpID string
	*rsa.PrivateKey
}

func (t *TwiplaJWTIssuer) New() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":     "twipla-3as-go-sdk",
		"roles":   []string{"intp"},
		"intp_id": t.intpID,
		"exp":     time.Now().Add(time.Hour * 4).Unix(),
	})
	token.Header["typ"] = "JWT"
	token.Header["kid"] = t.intpID
	return token.SignedString(t.PrivateKey)
}

func NewTwiplaJWTIssuer(intpID, privateKey string) (*TwiplaJWTIssuer, error) {
	pkey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return nil, err
	}

	return &TwiplaJWTIssuer{
		intpID:     intpID,
		PrivateKey: pkey,
	}, nil
}
