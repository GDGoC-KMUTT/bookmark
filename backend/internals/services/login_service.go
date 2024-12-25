package services

import (
	"backend/internals/db/models"
	"backend/internals/entities/common"
	"backend/internals/entities/payload"
	"backend/internals/repositories"
	"backend/internals/utils"
	services2 "backend/internals/utils/services"
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
	oauthSvc services2.OAuthService
	jwtSvc   services2.Jwt
}

func NewLoginService(userRepo repositories.UserRepository, oauthClient services2.OAuthService, jwtSvc services2.Jwt) LoginService {
	return &loginService{
		userRepo: userRepo,
		oauthSvc: oauthClient,
		jwtSvc:   jwtSvc,
	}
}

func (r *loginService) OAuthSetup(body *payload.OauthCallback) (*oidc.UserInfo, error) {
	// * exchange code for token
	token, err := r.oauthSvc.Exchange(context.Background(), *body.Code)
	if err != nil {
		return nil, err
	}

	// * parse ID token from OAuth2 token
	userInfo, err := r.oauthSvc.UserInfo(context.TODO(), oauth2.StaticTokenSource(token))
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
	jwtToken := r.jwtSvc.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwtToken, err3 := r.jwtSvc.SignedString(jwtToken, []byte(*secret))
	if err3 != nil {
		return nil, err3
	}

	return &signedJwtToken, nil
}
