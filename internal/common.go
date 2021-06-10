package internal

import (
	"errors"
	"github.com/spf13/viper"
	"time"

	"github.com/brianstrauch/spotify"
)

func Authenticate() (*spotify.API, error) {
	if time.Now().Unix() > viper.GetInt64("expiration") {
		if err := refresh(); err != nil {
			return nil, err
		}
	}

	token := viper.GetString("token")
	if token == "" {
		return nil, errors.New(NotLoggedInErr)
	}

	return spotify.NewAPI(token), nil
}

func refresh() error {
	refresh := viper.GetString("refresh_token")

	token, err := spotify.RefreshToken(refresh)
	if err != nil {
		return err
	}

	return SaveToken(token)
}

func SaveToken(token *spotify.Token) error {
	expiration := time.Now().Unix() + int64(token.ExpiresIn)

	viper.Set("expiration", expiration)
	viper.Set("token", token.AccessToken)
	viper.Set("refresh_token", token.RefreshToken)

	return viper.WriteConfig()
}

func WaitForUpdatedPlayback(api spotify.APIInterface, isUpdated func(playback *spotify.Playback) bool) (*spotify.Playback, error) {
	timeout := time.After(time.Second)
	tick := time.Tick(100 * time.Millisecond)

	for {
		select {
		case <-timeout:
			return nil, errors.New("request timed out")
		case <-tick:
			playback, err := api.GetPlayback()
			if err != nil {
				return nil, err
			}

			if isUpdated(playback) {
				return playback, nil
			}
		}
	}
}

func Search(api spotify.APIInterface, query string) (string, error) {
	page, err := api.Search(query, 1)
	if err != nil {
		return "", err
	}

	return page.Tracks.Items[0].URI, nil
}
