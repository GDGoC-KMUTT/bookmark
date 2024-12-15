package services

import "github.com/golang-jwt/jwt/v5"

// Jwt interface defines methods for creating and signing JWT tokens.
type Jwt interface {
	NewWithClaims(method jwt.SigningMethod, claims jwt.Claims) *jwt.Token
	SignedString(token *jwt.Token, secret interface{}) (string, error)
}
