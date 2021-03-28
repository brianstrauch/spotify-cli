package internal

import (
	"spotify/pkg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewPlayCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "play",
		Short:   "Play music.",
		PreRunE: checkToken,
		RunE: func(cmd *cobra.Command, _ []string) error {
			token := viper.GetString("token")
			return pkg.Play(token)
		},
	}
}
