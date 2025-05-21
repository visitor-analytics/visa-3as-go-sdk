package visa

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type tokenSigner struct {
	privateKey *rsa.PrivateKey
	intpID     string
}

func (t *tokenSigner) IntpToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":     "twipla-3as-go-sdk",
		"roles":   []string{"intp"},
		"intp_id": t.intpID,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 4).Unix(),
	})

	token.Header["kid"] = t.intpID
	return token.SignedString(t.privateKey)
}

func (t *tokenSigner) IntpcToken(intpcID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":      "twipla-3as-go-sdk",
		"roles":    []string{"intpc"},
		"intp_id":  t.intpID,
		"intpc_id": intpcID,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 4).Unix(),
	})

	token.Header["kid"] = t.intpID
	return token.SignedString(t.privateKey)
}

type authTransport struct {
	http.RoundTripper
	signer *tokenSigner
}

func (t *authTransport) RoundTrip(r *http.Request) (*http.Response, error) {

	return t.RoundTripper.RoundTrip(r)
}
