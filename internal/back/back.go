package back

import (
	"spotify/internal"
	"spotify/pkg"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "back",
		Short: "Skip to previous song.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			return back(api)
		},
	}
}

func back(api pkg.APIInterface) error {
	return api.Previous()
}
