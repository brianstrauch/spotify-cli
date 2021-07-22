package p

import (
	"errors"
	"github.com/spf13/cobra"
	"spotify/internal"
	"spotify/internal/pause"
	"spotify/internal/play"
	"strings"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "p [song]",
		// Keep hidden, since this command is an alias.
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}
			query := strings.Join(args, " ")
			queryType := "track"

			deviceID, err := cmd.Flags().GetString("device-id")
			if err != nil {
				return err
			}

			contextQuery, err := cmd.Flags().GetString("playlist")
			if err != nil {
				return err
			}

			if contextQuery == "" {
				contextQuery, err = cmd.Flags().GetString("album")
				if err != nil {
					return err
				}
				queryType = "album"
			} else {
				queryType = "playlist"
			}

			status, err := p(api, query, contextQuery, queryType, deviceID)
			if err != nil {
				return err
			}

			cmd.Print(status)
			return nil
		},
	}

	cmd.Flags().String("device-id", "", "device ID from 'spotify device list'")
	cmd.Flags().String("playlist", "", "playlist name from 'spotify playlist list'")
	cmd.Flags().String("album", "", "album name")

	return cmd
}

func p(api internal.APIInterface, query, contextQuery, queryType, deviceID string) (string, error) {
	if len(query) > 0 {
		return play.Play(api, query, "", "track", deviceID)
	}

	playback, err := api.GetPlayback()
	if err != nil {
		return "", err
	}

	if playback == nil {
		return "", errors.New(internal.ErrNoActiveDevice)
	}

	if playback.IsPlaying {
		return pause.Pause(api, deviceID)
	} else {
		return play.Play(api, query, contextQuery, queryType, deviceID)
	}
}
