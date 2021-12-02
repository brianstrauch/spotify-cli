package p

import (
	"errors"
	"github.com/spf13/cobra"
	"spotify/internal"
	"spotify/internal/pause"
	"spotify/internal/play"
	"spotify/internal/playlist"
	"strings"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "p [song]",
		Hidden: true, // Keep hidden, since this command is an alias.
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

			status, err := p(api, track, playlist, album)
			if err != nil {
				return err
			}

			cmd.Print(status)
			return nil
		},
	}

	cmd.Flags().String("playlist", "", "playlist name from 'spotify playlist list'")
	cmd.Flags().String("album", "", "album name")

	_ = cmd.RegisterFlagCompletionFunc("playlist", playlist.AutocompletePlaylist)

	return cmd
}

func p(api internal.APIInterface, track, playlist, album string) (string, error) {
	if track != "" || playlist != "" || album != "" {
		return play.Play(api, track, playlist, album)
	}

	playback, err := api.GetPlayback()
	if err != nil {
		return "", err
	}
	if playback == nil {
		return "", errors.New(internal.ErrNoActiveDevice)
	}

	if playback.IsPlaying {
		return pause.Pause(api)
	} else {
		return play.Play(api, "", "", "")
	}
}
