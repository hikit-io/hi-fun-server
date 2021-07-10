package services

import "golang.org/x/oauth2"

var (
	_OAuth2 = oauth2.Config{
		ClientID:     "auth_service",
		ClientSecret: "auth_service_secret",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://hfunc.nekilc.cn:8010/api/v1/auth_service/authorize",
			TokenURL: "https://hfunc.nekilc.cn:8010/api/v1/auth_service/token",
		},
		RedirectURL: "https://hfunc.nekilc.cn:8020/api/v1/user_service/oauth2",
		Scopes:      []string{"rw"},
	}
)

func OAuth2() *oauth2.Config {
	return &_OAuth2
}
