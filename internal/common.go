package internal

import (
	"errors"
	"time"

	"github.com/brianstrauch/spotify"
	"github.com/spf13/viper"
)

const ClientID = "81dddfee3e8d47d89b7902ba616f3357"

func Authenticate() (*spotify.API, error) {
	if time.Now().Unix() > viper.GetInt64("expiration") {
		if err := refresh(); err != nil {
			return nil, err
		}
	}

	token := viper.GetString("token")
	if token == "" {
		return nil, errors.New(ErrNotLoggedIn)
	}

	return spotify.NewAPI(token), nil
}

func refresh() error {
	refresh := viper.GetString("refresh_token")

	token, err := spotify.RefreshPKCEToken(refresh, ClientID)
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

func WaitForUpdatedPlayback(api APIInterface, isUpdated func(playback *spotify.Playback) bool) (*spotify.Playback, error) {
	timeout := time.After(time.Second)
	tick := time.NewTicker(100 * time.Millisecond)

	for {
		select {
		case <-timeout:
			return nil, errors.New("request timed out")
		case <-tick.C:
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

func Search(api APIInterface, query string) (*spotify.Track, error) {
	page, err := api.Search(query, 1)
	if err != nil {
		return nil, err
	}
	return page.Tracks.Items[0], nil
}
