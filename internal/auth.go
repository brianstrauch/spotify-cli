package internal

import (
	"errors"
	"spotify/internal/login"
	"spotify/pkg"
	"time"

	"github.com/spf13/viper"
)

func Authenticate() (*pkg.API, error) {
	if time.Now().Unix() > viper.GetInt64("expiration") {
		if err := refresh(); err != nil {
			return nil, err
		}
	}

	token := viper.GetString("token")
	if token == "" {
		return nil, errors.New(NotLoggedInErr)
	}

	return pkg.NewAPI(token), nil
}

func refresh() error {
	refresh := viper.GetString("refresh_token")

	token, err := pkg.RefreshToken(refresh)
	if err != nil {
		return err
	}

	return login.SaveToken(token)
}
