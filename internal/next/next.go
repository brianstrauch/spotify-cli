package next

import (
	"spotify/internal"
	"spotify/pkg"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "next",
		Short: "Skip to next song.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			return next(api)
		},
	}
}

func next(api pkg.APIInterface) error {
	return api.Next()
}
