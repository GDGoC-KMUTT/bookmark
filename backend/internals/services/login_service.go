package services

import (
	"backend/internals/db/models"
	"backend/internals/entities/common"
	"backend/internals/entities/payload"
	"backend/internals/repositories"
	"context"
	"errors"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
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
	fmt.Println("OIDC ID:", *oidcClaims.Id)
	user := new(models.User)
	if err := r.userRepo.First(user, "oid = ?", *oidcClaims.Id); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	fmt.Println("after first operation")
	// * if user not exist, create new user
	if user.Id == nil {
		fmt.Println("id is null")
		user = &models.User{
			Id:        nil,
			Oid:       oidcClaims.Id,
			Firstname: oidcClaims.FirstName,
			Lastname:  oidcClaims.Lastname,
			Email:     oidcClaims.Email,
			PhotoUrl:  oidcClaims.Picture,
			CreatedAt: nil,
			UpdatedAt: nil,
		}
		fmt.Println("creating user...")
		if err := r.userRepo.Create(user); err != nil {
			fmt.Println("failed to create user: ", err.Error())
			return nil, err
		}
	}

	return user, nil
}
