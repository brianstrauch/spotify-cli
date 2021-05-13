package play

import (
	"errors"
	"spotify/internal"
	"spotify/internal/status"
	"spotify/pkg"
	"spotify/pkg/model"
	"strings"

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

			var query string
			if len(args) > 0 {
				query = strings.Join(args, " ")
			}

			status, err := Play(api, query)
			if err != nil {
				return err
			}

			cmd.Print(status)
			return nil
		},
	}
}

func Play(api pkg.APIInterface, query string) (string, error) {
	playback, err := api.Status()
	if err != nil {
		return "", err
	}

	if playback == nil {
		return "", errors.New(internal.NoActiveDeviceErr)
	}

	var uri string

	if len(query) > 0 {
		uri, err = Search(api, query)
		if err != nil {
			return "", err
		}
	}

	if err := api.Play(uri); err != nil {
		if err.Error() == internal.RestrictionViolatedSpotifyErr {
			return "", errors.New(internal.AlreadyPlayingErr)
		}
	}

	playback, err = api.WaitForUpdatedPlayback(func(playback *model.Playback) bool {
		return len(playback.Item.ID) > 0 && playback.IsPlaying
	})
	if err != nil {
		return "", err
	}

	return status.Show(playback), nil
}

func Search(api pkg.APIInterface, query string) (string, error) {
	page, err := api.Search(query, 1)
	if err != nil {
		return "", err
	}

	return page.Tracks.Items[0].URI, nil
}
