package services

import (
	"backend/internals/db/models"
	"backend/internals/entities/payload"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type LoginService interface {
	OAuthSetup(body *payload.OauthCallback, oauthConfig *oauth2.Config, oidcProvider *oidc.Provider) (*oidc.UserInfo, error)
	GetOrCreateUserFromClaims(userInfo *oidc.UserInfo) (*models.User, error)
}
