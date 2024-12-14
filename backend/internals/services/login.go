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
	//Login(email *string, password *string) (accessToken *string, err error)
	//JwtSign(claims jwt.Claims, opts ...jwt.TokenOption) (string, error)
	//Me(userId *primitive.ObjectID) (user *entities.UserInfo, err error)
	//CreateUser(email *string, password *string) (NewUser *entities.UserInfo, err error)
	//CreateEmailVerifiedHash() (*string, *string, *time.Time, error)
	//SendEmailWithRateLimit(user *entities.UserInfo, subject string, templateFile string, token *string) error
	//VerifyEmail(userId *primitive.ObjectID, token *string) (*entities.UserInfo, error)
	//ResentEmailVerification(userId *primitive.ObjectID) (*entities.UserInfo, error)
	//AccountRecovery(email *string) (*entities.UserInfo, error)
	//NewPassword(userId *primitive.ObjectID, token *string, newPassword *string) (*entities.UserInfo, error)
}
