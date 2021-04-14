package pause

import (
	"errors"
	"spotify/internal"
	"spotify/pkg"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "pause",
		Short: "Pause music.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			return Pause(api)
		},
	}
}

func Pause(api pkg.APIInterface) error {
	err := api.Pause()

	if err != nil {
		switch err.Error() {
		case internal.RestrictionViolatedSpotifyErr:
			return errors.New(internal.AlreadyPausedErr)
		case internal.NoActiveDeviceSpotifyErr:
			return errors.New(internal.NoActiveDeviceErr)
		}
	}

	return err
}
