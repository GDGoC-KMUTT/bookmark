package utilServices

import (
	"backend/internals/config"
	"context"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"net/url"
)

type oauthService struct {
	config       *config.Config
	oauthClient  oauth2.Config
	oidcProvider *oidc.Provider
}

func NewOAuthService(conf *config.Config) OAuthService {
	authService := &oauthService{
		config: conf,
	}

	if conf.FrontendScheme != nil && conf.FrontendUrl != nil {
		frontendUrl := fmt.Sprintf("%s://%s", *conf.FrontendScheme, *conf.FrontendUrl)
		redirectUrl, err := url.JoinPath(frontendUrl, "/callback")
		if err != nil {
			logrus.Fatal("[OIDC] Unable to join url path", err)
		}

		authService.oidcProvider, err = oidc.NewProvider(context.Background(), *conf.OauthEndpoint)
		if err != nil {
			logrus.Fatal("[OIDC] Unable to create oidc provider", err)
		}

		authService.oauthClient = oauth2.Config{
			ClientID:     *conf.OauthClientId,
			ClientSecret: *conf.OauthClientSecret,
			RedirectURL:  redirectUrl,
			Endpoint:     authService.oidcProvider.Endpoint(),
			Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
		}
		return authService
	}
	return authService
}

func (r *oauthService) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := r.oauthClient.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (r *oauthService) UserInfo(ctx context.Context, tokenSource oauth2.TokenSource) (*oidc.UserInfo, error) {
	userInfo, err := r.oidcProvider.UserInfo(ctx, tokenSource)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
