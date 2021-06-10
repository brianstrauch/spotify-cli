package back

import (
	"errors"
	"spotify/internal"
	"spotify/internal/status"

	"github.com/brianstrauch/spotify"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "back",
		Aliases: []string{"b"},
		Short:   "Skip to previous song.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			status, err := back(api)
			if err != nil {
				return err
			}

			cmd.Print(status)
			return nil
		},
	}
}

func back(api spotify.APIInterface) (string, error) {
	playback, err := api.GetPlayback()
	if err != nil {
		return "", err
	}

	if playback == nil {
		return "", errors.New(internal.NoActiveDeviceErr)
	}

	id := playback.Item.ID

	if err := api.SkipToPreviousTrack(); err != nil {
		if err.Error() == internal.RestrictionViolatedSpotifyErr {
			return "", errors.New(internal.NoPreviousErr)
		}
	}

	playback, err = internal.WaitForUpdatedPlayback(api, func(playback *spotify.Playback) bool {
		return playback.Item.ID != id
	})
	if err != nil {
		return "", err
	}

	return status.Show(playback), nil
}
