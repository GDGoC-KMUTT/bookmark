package controllers

import (
	"backend/internals/config"
	"backend/internals/db/models"
	"backend/internals/entities/common"
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/utils"
	"context"
	"errors"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
	"net/url"
)

type LoginController struct {
	config       *config.Config
	database     *gorm.DB
	OidcProvider *oidc.Provider
	OidcVerifier *oidc.IDTokenVerifier
	Oauth2Config *oauth2.Config
}

func NewLoginController(config *config.Config, database *gorm.DB) LoginController {
	controller := LoginController{
		config:   config,
		database: database,
	}

	redirectUrl, err := url.JoinPath(*config.FrontendUrl, "/callback")
	if err != nil {
		logrus.Fatal("[OIDC] Unable to join url path", err)
	}

	controller.OidcProvider, err = oidc.NewProvider(context.Background(), *config.OauthEndpoint)
	if err != nil {
		logrus.Fatal("[OIDC] Unable to create oidc provider", err)
	}

	controller.Oauth2Config = &oauth2.Config{
		ClientID:     *config.OauthClientId,
		ClientSecret: *config.OauthClientSecret,
		RedirectURL:  redirectUrl,
		Endpoint:     controller.OidcProvider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	return controller
}

func (r *LoginController) LoginRedirect(c *fiber.Ctx) error {
	return c.Redirect(r.Oauth2Config.AuthCodeURL("state"))
}

func (r *LoginController) LoginCallBack(c *fiber.Ctx) error {
	// * parse body
	body := new(payload.OauthCallback)
	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "unable to parse body",
		}
	}

	// * validate body
	if err := utils.Validate.Struct(body); err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		return &response.GenericError{
			Err:     validationErrors,
			Message: "invalid body",
		}
	}

	// * exchange code for token
	token, err := r.Oauth2Config.Exchange(context.Background(), *body.Code)
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to exchange code for token",
		}
	}

	// * parse ID token from OAuth2 token
	userInfo, err := r.OidcProvider.UserInfo(context.TODO(), oauth2.StaticTokenSource(token))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to get user info",
		}
	}

	// * parse user claims
	oidcClaims := new(common.OIdClaims)
	if err := userInfo.Claims(oidcClaims); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to parse user claims",
		}
	}

	// * first user with oid
	user := new(models.User)
	if tx := r.database.First(user, "oid = ?", oidcClaims.Id); tx.Error != nil {
		if tx.Error.Error() != "record not found" {
			return &response.GenericError{
				Err:     err,
				Message: "failed to query user",
			}
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
			CreatedAt: nil,
			UpdatedAt: nil,
		}

		if tx := r.database.Create(user); tx.Error != nil {
			return &response.GenericError{
				Err:     err,
				Message: "failed to create user",
			}
		}
	}

	// * generate jwt token
	claims := jwt.MapClaims{
		"userId": user.Id,
	}

	// Sign JWT token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwtToken, err := jwtToken.SignedString([]byte(*r.config.SecretKey))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to sign jwt token",
		}
	}

	// * set cookie
	c.Cookie(&fiber.Cookie{
		Name:  "login",
		Value: signedJwtToken,
	})

	result := new(payload.CallbackResponse)
	result.Token = &signedJwtToken

	return response.Ok(c, result)
}
