package auth

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	goauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type Auth struct {
	oauth2Config *oauth2.Config
}

func NewAuth(config Config) *Auth {
	a := &Auth{}

	c, err := google.ConfigFromJSON(config.creds, "https://www.googleapis.com/auth/userinfo.email")
	if err != nil {
		panic(err.Error())
	}

	a.oauth2Config = c

	return a
}

func (a *Auth) GetLoginURL() string {
	return a.oauth2Config.AuthCodeURL("state")
}

func (a *Auth) Exchange(ctx context.Context, code string) oauth2.TokenSource {
	token, err := a.oauth2Config.Exchange(ctx, code)
	if err != nil {
		panic(err.Error())
	}

	return a.oauth2Config.TokenSource(ctx, token)
}

func (a *Auth) ExchangeForGID(ctx context.Context, code string) string {
	ts := a.Exchange(ctx, code)
	return a.GetGID(ctx, ts)
}

func (a *Auth) GetGID(ctx context.Context, ts oauth2.TokenSource) string {
	s, err := goauth2.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		panic(err.Error())
	}

	srv := goauth2.NewUserinfoV2Service(s)

	userinfo, err := srv.Me.Get().Do()
	if err != nil {
		panic(err.Error())
	}

	return userinfo.Id
}
