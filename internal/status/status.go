package status

import (
	"errors"
	"spotify/internal"
	"spotify/internal/login"
	"spotify/pkg"
	"spotify/pkg/model"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show the current song.",
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
			playback, err := status(api)
			if err != nil {
				return err
			}

			cmd.Println(playback.Item.Name)
			for _, artist := range playback.Item.Artists {
				cmd.Println(artist.Name)
			}
			return nil
		},
	}
}

func status(api pkg.APIInterface) (*model.Playback, error) {
	return api.Status()
}
