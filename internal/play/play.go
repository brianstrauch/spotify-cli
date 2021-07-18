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
	cmd := &cobra.Command{
		Use:   "play [song]",
		Short: "play current song, or a specific song",
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			query := strings.Join(args, " ")

			deviceID, err := cmd.Flags().GetString("device-id")
			if err != nil {
				return err
			}

			status, err := Play(api, query, deviceID)
			if err != nil {
				return err
			}

			cmd.Print(status)
			return nil
		},
	}

	cmd.Flags().String("device-id", "", "device ID from 'spotify device list'")

	return cmd
}

func Play(api internal.APIInterface, query, deviceID string) (string, error) {
	playback, err := api.GetPlayback()
	if err != nil {
		return "", err
	}

	if playback == nil {
		return "", errors.New(internal.ErrNoActiveDevice)
	}

	if len(query) > 0 {
		track, err := internal.Search(api, query)
		if err != nil {
			return "", err
		}

		if err := api.Play(deviceID, track.URI); err != nil {
			return "", err
		}
	} else {
		if err := api.Play(deviceID); err != nil {
			return "", err
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
