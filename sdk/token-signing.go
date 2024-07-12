package sdk

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
)

type TokenSigner interface {
	Sign(kid string, payload jwt.MapClaims) (string, error)
}

type RS256 struct {
	PrivateKey *rsa.PrivateKey
}

func (t *RS256) Sign(kid string, payload jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, payload)

	token.Header["kid"] = kid
	token.Header["typ"] = "jwt"
	token.Header["alg"] = "RS256"

	return token.SignedString(t.PrivateKey)
}
