package play

import (
	"errors"
	"spotify/internal"
	"spotify/internal/status"
	"strings"

	"github.com/brianstrauch/spotify"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "play [song]",
		Short: "Play current song, or a specific song.",
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			query := strings.Join(args, " ")

			status, err := Play(api, query)
			if err != nil {
				return err
			}

			cmd.Print(status)
			return nil
		},
	}
}

func Play(api spotify.APIInterface, query string) (string, error) {
	playback, err := api.GetPlayback()
	if err != nil {
		return "", err
	}

	if playback == nil {
		return "", errors.New(internal.NoActiveDeviceErr)
	}

	if len(query) > 0 {
		uri, err := internal.Search(api, query)
		if err != nil {
			return "", err
		}

		err = api.Play(uri)
	} else {
		err = api.Play()
	}

	if err != nil {
		if err.Error() == internal.RestrictionViolatedSpotifyErr {
			return "", errors.New(internal.AlreadyPlayingErr)
		}
	}

	playback, err = internal.WaitForUpdatedPlayback(api, func(playback *spotify.Playback) bool {
		// The first check safeguards against empty playback objects
		return len(playback.Item.ID) > 0 && playback.IsPlaying
	})
	if err != nil {
		return "", err
	}

	return status.Show(playback), nil
}
