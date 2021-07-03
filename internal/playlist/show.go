package playlist

import (
	"errors"
	"fmt"
	"spotify/internal"
	"strings"

	"github.com/brianstrauch/spotify"
	"github.com/spf13/cobra"
)

func NewShowCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show artist and songs.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			name := strings.Join(args, " ")

			return Show(api, name)
		},
	}
}

func Show(api *spotify.API, name string) error {
	playlists, err := api.GetPlaylists()
	if err != nil {
		return err
	}

	id := ""
	for _, playlist := range playlists {
		if strings.EqualFold(playlist.Name, name) {
			id = playlist.ID
		}
	}
	if id == "" {
		return errors.New("no such playlist")
	}

	playlist, err := api.GetPlaylist(id)
	if err != nil {
		return err
	}

	output, err := formatPlaylist(api, playlist)
	if err != nil {
		return err
	}

	fmt.Print(output)
	return nil
}

func formatPlaylist(api *spotify.API, playlist *spotify.Playlist) (string, error) {
	list := fmt.Sprintf("💿 %s\n", playlist.Name)

	for i, track := range playlist.Tracks.Items {
		artists := make([]string, len(track.Track.Artists))
		for j, artist := range track.Track.Artists {
			if err := artist.Get(api, &artist); err != nil {
				return "", err
			}
			artists[j] = artist.Name
		}

		list += fmt.Sprintf("%d. %s - %s\n", i+1, track.Track.Name, strings.Join(artists, ", "))
	}

	return list, nil
}
