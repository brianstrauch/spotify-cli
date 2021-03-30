package internal

import (
	"errors"
	"spotify/pkg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewPauseCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "pause",
		Short: "Pause music.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			token := new(pkg.Token)
			if err := viper.UnmarshalKey("token", token); err != nil {
				return errors.New(NotLoggedInErr)
			}

			api := pkg.NewAPI(token)
			return pause(api)
		},
	}
}

func pause(api pkg.APIInterface) error {
	return api.Pause()
}
