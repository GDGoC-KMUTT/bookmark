package utilServices

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type OAuthService interface {
	// Exchange OAuth2Client
	Exchange(ctx context.Context, code string) (*oauth2.Token, error)
	// UserInfo OIDCProvider
	UserInfo(ctx context.Context, tokenSource oauth2.TokenSource) (*oidc.UserInfo, error)
}
