package utils

import (
	"crypto/rsa"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func DecodePublicKey(publicKey string) (*rsa.PublicKey, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		return nil, errors.New("failed to parse public key")
	}
	return key, nil
}
