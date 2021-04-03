package internal

import (
	"errors"
	"spotify/pkg"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewPlayCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "play",
		Short: "Play music.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			token := viper.GetString("token")
			if token == "" {
				return errors.New(NotLoggedInErr)
			}

			if time.Now().Unix() > viper.GetInt64("expiration") {
				return errors.New(TokenExpiredErr)
			}

			api := pkg.NewAPI(token)
			return play(api)
		},
	}
}

func play(api pkg.APIInterface) error {
	return api.Play()
}
