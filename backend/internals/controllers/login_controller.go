package controllers

import (
	"backend/internals/config"
	"backend/internals/entities/common"
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	"backend/internals/utils"
	"context"
	"errors"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"net/url"
)

type LoginController struct {
	config       *config.Config
	OidcProvider *oidc.Provider
	OidcVerifier *oidc.IDTokenVerifier
	Oauth2Config *oauth2.Config
	loginSvc     services.LoginService
}

func NewLoginController(config *config.Config, loginSvc services.LoginService) LoginController {
	controller := LoginController{
		config:   config,
		loginSvc: loginSvc,
	}

	frontendUrl := fmt.Sprintf("%s://%s", *config.FrontendScheme, *config.FrontendUrl)
	redirectUrl, err := url.JoinPath(frontendUrl, "/callback")
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

// LoginRedirect
// @ID loginRedirect
// @Tags login
// @Summary LoginRedirect
// @Failure 400 {object} response.GenericError
// @Router /login/redirect [get]
func (r *LoginController) LoginRedirect(c *fiber.Ctx) error {
	return c.Redirect(r.Oauth2Config.AuthCodeURL("state"))
}

// LoginCallBack
// @ID loginCallBack
// @Tags login
// @Summary LoginCallBack
// @Param q body payload.OauthCallback true "OauthCallback"
// @Accept json
// @Produce json
// @Success 200 {object} response.InfoResponse[payload.CallbackResponse]
// @Failure 400 {object} response.GenericError
// @Router /login/callback [post]
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

	userInfo, err := r.loginSvc.OAuthSetup(body, r.Oauth2Config, r.OidcProvider)
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to setup OAuth",
		}
	}

	// * parse user claims
	oidcClaims := new(common.OIdClaims)
	if err1 := userInfo.Claims(oidcClaims); err1 != nil {
		return &response.GenericError{
			Err:     err1,
			Message: "failed to parse user claims",
		}
	}

	user, err2 := r.loginSvc.GetOrCreateUserFromClaims(userInfo)
	if err2 != nil {
		return &response.GenericError{
			Err:     err2,
			Message: "failed to get or create user from claims",
		}
	}

	// * generate jwt token
	claims := jwt.MapClaims{
		"userId": user.Id,
	}

	// Sign JWT token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwtToken, err3 := jwtToken.SignedString([]byte(*r.config.SecretKey))
	if err3 != nil {
		return &response.GenericError{
			Err:     err3,
			Message: "failed to sign jwt token",
		}
	}

	// * set cookie
	c.Cookie(&fiber.Cookie{
		Name:  "login",
		Value: signedJwtToken,
	})

	return response.Ok(c, &payload.CallbackResponse{
		Token: &signedJwtToken,
	})
}
