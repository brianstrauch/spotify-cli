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
		Short: "Show playlist artist and songs.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}
			return Show(api, args)
		},
	}
}

func Show(api *spotify.API, args []string) error {
	playlists, err := api.GetPlaylists()
	if err != nil {
		return err
	}
	id := ""
	for _, playlist := range playlists {
		if strings.EqualFold(strings.ToLower(playlist.Name), strings.ToLower(args[0])) {
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
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("📝: %s\n", playlist.Name))
	builder.WriteString(fmt.Sprintln("💿 Tracks:"))
	for i, track := range playlist.Tracks.Items {
		artistNames := make([]string, len(track.Track.Artists))
		for i, artist := range track.Track.Artists {
			if err := artist.Get(api, &artist); err != nil {
				return "", err
			}
			artistNames[i] = artist.Name
		}
		builder.WriteString(
			fmt.Sprintf(
				"%d. %s - %s\n",
				i+1,
				strings.Join(artistNames, ". "),
				track.Track.Name,
			),
		)
	}
	return builder.String(), nil
}
