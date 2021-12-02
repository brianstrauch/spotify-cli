package play

import (
	"errors"
	"spotify/internal"
	"spotify/internal/playlist"
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

			track := strings.Join(args, " ")

			playlist, err := cmd.Flags().GetString("playlist")
			if err != nil {
				return err
			}

			album, err := cmd.Flags().GetString("album")
			if err != nil {
				return err
			}

			if track != "" && playlist != "" || track != "" && album != "" || playlist != "" && album != "" {
				return errors.New(internal.ErrInvalidPlayArgs)
			}

			status, err := Play(api, track, playlist, album)
			if err != nil {
				return err
			}

			cmd.Print(status)
			return nil
		},
	}

	cmd.Flags().String("playlist", "", "playlist name from 'spotify playlist list'")
	cmd.Flags().String("album", "", "album name that you wish to play")

	_ = cmd.RegisterFlagCompletionFunc("playlist", playlist.AutocompletePlaylist)

	return cmd
}

func Play(api internal.APIInterface, track, playlist, album string) (string, error) {
	playback, err := api.GetPlayback()
	if err != nil {
		return "", err
	}
	if playback == nil {
		return "", errors.New(internal.ErrNoActiveDevice)
	}

	isPlaying := playback.IsPlaying
	id := playback.Item.ID
	progressMs := playback.ProgressMs

	if track != "" {
		track, err := internal.SearchTrack(api, track)
		if err != nil {
			return "", err
		}

		if err := api.Play("", track.URI); err != nil {
			return "", err
		}
	} else if album != "" {
		album, err := internal.SearchAlbum(api, album)
		if err != nil {
			return "", err
		}

		if err := api.Play(album.URI); err != nil {
			return "", err
		}
	} else if playlist != "" {
		playlist, err := internal.SearchPlaylist(api, playlist)
		if err != nil {
			return "", err
		}

		if err := api.Play(playlist.URI); err != nil {
			return "", err
		}
	} else {
		if err := api.Play(""); err != nil {
			return "", err
		}
	}

	playback, err = internal.WaitForUpdatedPlayback(api, func(playback *spotify.Playback) bool {
		hasChanged := playback.Item.ID != "" && (playback.Item.ID != id || playback.ProgressMs < progressMs)
		return !isPlaying && playback.IsPlaying || hasChanged
	})
	if err != nil {
		return "", err
	}

	return status.Show(playback), nil
}
