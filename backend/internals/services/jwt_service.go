package services

import "github.com/golang-jwt/jwt/v5"

// jwtService implements the Jwt interface.
type jwtService struct{}

func NewJwtService() Jwt {
	return &jwtService{}
}

// NewWithClaims creates a new JWT token with the given signing method and claims.
func (j jwtService) NewWithClaims(method jwt.SigningMethod, claims jwt.Claims) *jwt.Token {
	return jwt.NewWithClaims(method, claims)
}

// SignedString signs the token using the given secret and returns the signed string.
func (j jwtService) SignedString(token *jwt.Token, secret interface{}) (string, error) {
	return token.SignedString(secret)
}
