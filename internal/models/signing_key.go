package models

import (
	"crypto/rsa"
	"encoding/base64"
	"math/big"
)

type SigningKey struct {
	ID         string
	PrivateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
}

func (s *SigningKey) PublicKeyToJWK() map[string]string {
	n := base64.RawURLEncoding.EncodeToString(s.PublicKey.N.Bytes())
    e := base64.RawURLEncoding.EncodeToString(big.NewInt(int64(s.PublicKey.E)).Bytes())

    return map[string]string{
        "kty": "RSA",
        "kid": s.ID,
        "use": "sig",
        "alg": "RS256",
        "n":   n,
        "e":   e,
    }
}
