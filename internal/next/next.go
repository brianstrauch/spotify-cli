package next

import (
	"errors"
	"spotify/internal"
	"spotify/internal/status"

	"github.com/brianstrauch/spotify"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "next",
		Aliases: []string{"n"},
		Short:   "skip to next song",
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			status, err := next(api)
			if err != nil {
				return err
			}

			cmd.Print(status)
			return nil
		},
	}
}

func next(api internal.APIInterface) (string, error) {
	playback, err := api.GetPlayback()
	if err != nil {
		return "", err
	}

	if playback == nil {
		return "", errors.New(internal.ErrNoActiveDevice)
	}

	id := playback.Item.ID
	progressMs := playback.ProgressMs

	if err := api.SkipToNextTrack(); err != nil {
		return "", err
	}

	playback, err = internal.WaitForUpdatedPlayback(api, func(playback *spotify.Playback) bool {
		return playback.Item.ID != id || playback.ProgressMs < progressMs
	})
	if err != nil {
		return "", err
	}

	return status.Show(playback), nil
}
