package player

import (
	"os"

	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/gateway/repository"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type InteractorI interface {
	CreatePlayer(isNewPlayer string) string
	SetInactive(playerId string)
	OauthLogin() string
	GoogleCallback(state, code string) ([]byte, error)
}

func NewService(repo repository.PlayerRepositoryI) InteractorI {
	return &Service{
		repo: repo,
		googleOauthConfig: &oauth2.Config{
			RedirectURL:  "http://localhost:1323/callback",
			ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		},
	}
}

type Service struct {
	repo              repository.PlayerRepositoryI
	googleOauthConfig *oauth2.Config
}
