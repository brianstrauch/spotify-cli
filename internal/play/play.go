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

			status, err := Play(api, query, contextQuery, queryType, deviceID)
			if err != nil {
				return err
			}

			cmd.Print(status)
			return nil
		},
	}

	cmd.Flags().String("device-id", "", "device ID from 'spotify device list'")
	cmd.Flags().String("playlist", "", "playlist name from 'spotify playlist list'")
	cmd.Flags().String("album", "", "album name that you wish to play")

	return cmd
}

func Play(api internal.APIInterface, query, contextQuery, queryType, deviceID string) (string, error) {
	playback, err := api.GetPlayback()
	if err != nil {
		return "", err
	}

	if playback == nil {
		return "", errors.New(internal.ErrNoActiveDevice)
	}

	switch queryType {
	case "album":
		api, err := internal.Authenticate()
		if err != nil {
			return "", err
		}

		paging, err := api.Search(contextQuery, "album", 1)
		if err != nil {
			return "", err
		}

		albums := paging.Albums.Items
		//TODO: Implement Checking of no returns of albums

		if err := api.Play(deviceID, albums[0].URI); err != nil {
			return "", err
		}

	case "playlist":
		// Return a different API interface required for the playlist commands?
		api, err := internal.Authenticate()
		if err != nil {
			return "", err
		}

		playlists, err := api.GetPlaylists()
		if err != nil {
			return "", err
		}

		for _, playlist := range playlists {
			if strings.EqualFold(playlist.Name, contextQuery) {
				if err := api.Play(deviceID, playlist.URI); err != nil {
					return "", err
				}
				break
			}
		}

	default:
		track, err := internal.Search(api, query, "track")
		if err != nil {
			return "", err
		}

		if err := api.Play(deviceID, "", track.URI); err != nil {
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
