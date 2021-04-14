package next

import (
	"errors"
	"spotify/internal"
	"spotify/pkg"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "next",
		Aliases: []string{"n"},
		Short:   "Skip to next song.",
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
	err := api.Next()

	if err != nil {
		switch err.Error() {
		case internal.RestrictionViolatedSpotifyErr:
			return errors.New(internal.NoNextErr)
		case internal.NoActiveDeviceSpotifyErr:
			return errors.New(internal.NoActiveDeviceErr)
		}
	}

	return err
}
