package internal

import (
	"errors"
	"spotify/pkg"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewPauseCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "pause",
		Short: "Pause music.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			token := viper.GetString("token")
			if token == "" {
				return errors.New(NotLoggedInErr)
			}

			if time.Now().Unix() > viper.GetInt64("expiration") {
				return errors.New(TokenExpiredErr)
			}

			api := pkg.NewAPI(token)
			return pause(api)
		},
	}
}

func pause(api pkg.APIInterface) error {
	return api.Pause()
}
