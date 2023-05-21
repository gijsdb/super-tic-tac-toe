package player

import (
	"os"

	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/gateway/repository"
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type InteractorI interface {
	Create() string
	Get(id string) (*entity.Player, error)
	GetByEmail(email string) (*entity.Player, error)
	Update(player *entity.Player) *entity.Player
	OauthLogin(temp_player_id string) string
	GoogleCallback(state, code string) (string, error)
}

func NewService(repo repository.PlayerRepositoryI) InteractorI {
	return &PlayerService{
		repo: repo,
		googleOauthConfig: &oauth2.Config{
			RedirectURL:  "http://localhost:1323/callback",
			ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		},
		oauthStateStrings: make(map[string]*OauthStateString),
	}
}

type PlayerService struct {
	repo              repository.PlayerRepositoryI
	googleOauthConfig *oauth2.Config
	oauthStateStrings map[string]*OauthStateString
}

// Todo move to entities?
type OauthStateString struct {
	Active         bool
	Temp_player_id string
}
