package player

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

var (
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func (s *Service) OauthLogin() string {
	url := s.googleOauthConfig.AuthCodeURL(oauthStateString)

	return url
}

func (s *Service) GoogleCallback(state, code string) ([]byte, error) {
	content, err := getUserInfo(state, code, s.googleOauthConfig)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Printf("Content: %s\n", content)
	return content, nil
}

func getUserInfo(state string, code string, conf *oauth2.Config) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}
