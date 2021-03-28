package internal

import (
	"spotify/pkg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewPauseCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "pause",
		Short:   "Pause music.",
		PreRunE: checkToken,
		RunE: func(cmd *cobra.Command, _ []string) error {
			token := viper.GetString("token")
			return pkg.Pause(token)
		},
	}
}
