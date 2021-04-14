package back

import (
	"errors"
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
	err := api.Back()

	if err != nil {
		switch err.Error() {
		case internal.RestrictionViolatedSpotifyErr:
			return errors.New(internal.NoPreviousErr)
		case internal.NoActiveDeviceSpotifyErr:
			return errors.New(internal.NoActiveDeviceErr)
		}
	}

	return err
}
