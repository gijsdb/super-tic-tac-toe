package player

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/gijsdb/super-tic-tac-toe/internal/entity"
)

func (s *PlayerService) OauthLogin(temp_player_id string) string {
	oauthStateString := generateStateString(charset)
	s.oauthStateStrings[oauthStateString] = &OauthStateString{
		Active:         true,
		Temp_player_id: temp_player_id,
	}

	url := s.googleOauthConfig.AuthCodeURL(oauthStateString)
	return url
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func generateStateString(charset string) string {
	b := make([]byte, 32)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (s *PlayerService) GoogleCallback(state, code string) (string, error) {
	content, state_string, err := s.getUserInfo(state, code)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	fmt.Printf("Content: %s\n", content)

	type oauthContent struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Picture       string `json:"picture"`
	}

	var temp oauthContent

	err = json.Unmarshal(content, &temp)
	if err != nil {
		return "", err
	}

	player := s.repo.Update(&entity.Player{
		ID:       state_string.Temp_player_id,
		GoogleID: temp.ID,
		Email:    temp.Email,
		Picture:  temp.Picture,
	})

	return player.ID, nil
}

func (s *PlayerService) getUserInfo(state string, code string) ([]byte, *OauthStateString, error) {
	state_string, ok := s.oauthStateStrings[state]
	if !ok {
		return nil, nil, fmt.Errorf("invalid state strings")
	}
	token, err := s.googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	delete(s.oauthStateStrings, state)
	return contents, state_string, nil
}
