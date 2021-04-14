package play

import (
	"errors"
	"spotify/internal"
	"spotify/pkg"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "play",
		Short: "Play music.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			return Play(api)
		},
	}
}

func Play(api pkg.APIInterface) error {
	err := api.Play()

	if err != nil {
		switch err.Error() {
		case internal.RestrictionViolatedSpotifyErr:
			return errors.New(internal.AlreadyPlayingErr)
		case internal.NoActiveDeviceSpotifyErr:
			return errors.New(internal.NoActiveDeviceErr)
		}
	}

	return err
}
