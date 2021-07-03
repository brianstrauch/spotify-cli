package playlist

import (
	"errors"
	"fmt"
	"spotify/internal"
	"strconv"
	"strings"

	"github.com/brianstrauch/spotify"
	"github.com/olekukonko/tablewriter"
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

	output, err := formatPlaylist(playlist)
	if err != nil {
		return err
	}

	fmt.Print(output)
	return nil
}

func formatPlaylist(playlist *spotify.Playlist) (string, error) {
	output := new(strings.Builder)

	table := tablewriter.NewWriter(output)
	table.SetBorder(false)

	table.SetHeader([]string{"#", "Title", "Artist(s)"})

	for i, playlistTrack := range playlist.Tracks.Items {
		track := playlistTrack.Track

		artists := make([]string, len(track.Artists))
		for j, artist := range track.Artists {
			artists[j] = artist.Name
		}

		table.Append([]string{strconv.Itoa(i + 1), track.Name, strings.Join(artists, ", ")})
	}
	table.Render()

	return output.String(), nil
}
