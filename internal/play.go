package internal

import (
	"errors"
	"spotify/pkg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewPlayCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "play",
		Short: "Play music.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			token := new(pkg.Token)
			if err := viper.UnmarshalKey("token", token); err != nil {
				return errors.New(NotLoggedInErr)
			}

			api := pkg.NewAPI(token)
			return play(api)
		},
	}
}

func play(api pkg.APIInterface) error {
	return api.Play()
}
