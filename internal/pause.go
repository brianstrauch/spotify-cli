package internal

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"spotify/pkg"
)

func NewPauseCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "pause",
		Short: "Pause music.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			token := viper.GetString("token")
			// TODO: Verify that token is up-to-date
			return pkg.Pause(token)
		},
	}
}
