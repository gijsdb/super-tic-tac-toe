package player

import (
	"os"

	"github.com/gijsdb/super-tic-tac-toe/internal/adapter/repository"
	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type ServiceI interface {
	Create() string
	Get(id string) (*entity.Player, error)
	GetByEmail(email string) (*entity.Player, error)
	Update(player *entity.Player) *entity.Player
	OauthLogin(temp_player_id string) string
	GoogleCallback(state, code string) (string, error)
	LoadDBPlayersIntoMemory()
	GetHighscores() ([]*entity.Player, error)
	SeedPlayers()
}

type PlayerService struct {
	memorystore       repository.PlayerMemoryRepoI
	database          repository.PlayerDatabaseRepoI
	googleOauthConfig *oauth2.Config
	oauthStateStrings map[string]*OauthStateString
}

type OauthStateString struct {
	Active         bool
	Temp_player_id string
}

func NewPlayerService(ms repository.PlayerMemoryRepoI, db repository.PlayerDatabaseRepoI) ServiceI {
	return &PlayerService{
		memorystore: ms,
		database:    db,
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
