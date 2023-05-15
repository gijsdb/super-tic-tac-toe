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
	GoogleCallback(state, code string) (string, error)
}

func NewService(playerRepo repository.PlayerRepositoryI) InteractorI {
	return &Service{
		playerRepo: playerRepo,
		googleOauthConfig: &oauth2.Config{
			RedirectURL:  "http://localhost:1323/callback",
			ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		},
		oauthStateStrings: make(map[string]bool),
	}
}

type Service struct {
	playerRepo        repository.PlayerRepositoryI
	googleOauthConfig *oauth2.Config
	oauthStateStrings map[string]bool
}
