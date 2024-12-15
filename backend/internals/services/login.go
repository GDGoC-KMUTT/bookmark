package services

import (
	"backend/internals/db/models"
	"backend/internals/entities/payload"
	"github.com/coreos/go-oidc/v3/oidc"
)

type LoginService interface {
	OAuthSetup(body *payload.OauthCallback) (*oidc.UserInfo, error)
	GetOrCreateUserFromClaims(userInfo *oidc.UserInfo) (*models.User, error)
	SignJwtToken(user *models.User, secret *string) (*string, error)
}
