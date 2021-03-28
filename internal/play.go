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
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			if !IsAuthenticated() {
				return errors.New("You are not logged in. Please use 'spotify login' before running this command.")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			token := viper.GetString("token")
			return pkg.Play(token)
		},
	}
}
