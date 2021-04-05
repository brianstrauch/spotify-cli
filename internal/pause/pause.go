package pause

import (
	"errors"
	"spotify/internal"
	"spotify/internal/login"
	"spotify/pkg"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "pause",
		Short: "Pause music.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if time.Now().Unix() > viper.GetInt64("expiration") {
				refreshToken := viper.GetString("refresh_token")

				token, err := pkg.RefreshToken(refreshToken)
				if err != nil {
					return err
				}

				if err := login.SaveToken(token); err != nil {
					return err
				}
			}

			token := viper.GetString("token")
			if token == "" {
				return errors.New(internal.NotLoggedInErr)
			}

			api := pkg.NewAPI(token)
			return pause(api)
		},
	}
}

func pause(api pkg.APIInterface) error {
	return api.Pause()
}
