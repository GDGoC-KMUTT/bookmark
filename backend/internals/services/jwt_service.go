package services

import "github.com/golang-jwt/jwt/v5"

type Jwt interface {
	NewWithClaims(method jwt.SigningMethod, claims jwt.Claims, opts ...jwt.TokenOption) *jwt.Token
}

type JwtToken interface {
	SignedString(key []byte) (string, error)
}
