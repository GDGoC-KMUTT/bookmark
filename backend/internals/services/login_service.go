package services

import (
	"backend/internals/db/models"
	"backend/internals/entities/common"
	"backend/internals/entities/payload"
	"backend/internals/repositories"
	"backend/internals/utils"
	"context"
	"errors"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
	"time"
)

type loginService struct {
	userRepo repositories.UserRepository
}

func NewLoginService(userRepo repositories.UserRepository) LoginService {
	return &loginService{
		userRepo: userRepo,
	}
}

func (r *loginService) OAuthSetup(body *payload.OauthCallback, oauthConfig *oauth2.Config, oidcProvider *oidc.Provider) (*oidc.UserInfo, error) {
	// * exchange code for token
	token, err := oauthConfig.Exchange(context.Background(), *body.Code)
	if err != nil {
		return nil, err
	}

	// * parse ID token from OAuth2 token
	userInfo, err := oidcProvider.UserInfo(context.TODO(), oauth2.StaticTokenSource(token))
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

func (r *loginService) GetOrCreateUserFromClaims(userInfo *oidc.UserInfo) (*models.User, error) {
	// * parse user claims
	oidcClaims := new(common.OIdClaims)
	if err := userInfo.Claims(oidcClaims); err != nil {
		return nil, err
	}

	// * first user with oid
	user, err := r.userRepo.FindFirstUserByOid(oidcClaims.Id)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	// * if user not exist, create new user
	if user.Id == nil {
		user = &models.User{
			Id:        nil,
			Oid:       oidcClaims.Id,
			Firstname: oidcClaims.FirstName,
			Lastname:  oidcClaims.Lastname,
			Email:     oidcClaims.Email,
			PhotoUrl:  oidcClaims.Picture,
			CreatedAt: utils.Ptr(time.Now()),
			UpdatedAt: utils.Ptr(time.Now()),
		}
		if err := r.userRepo.CreateUser(user); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (r *loginService) SignJwtToken(user *models.User, secret *string) (*string, error) {
	// * generate jwt token
	claims := jwt.MapClaims{
		"userId": user.Id,
	}

	// Sign JWT token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwtToken, err3 := jwtToken.SignedString([]byte(*secret))
	if err3 != nil {
		return nil, err3
	}

	return &signedJwtToken, nil

}
