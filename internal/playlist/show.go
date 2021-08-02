package playlist

import (
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
		Use:   "show [playlist]",
		Short: "show artist and songs",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			playlist := strings.Join(args, " ")

			return show(api, playlist)
		},
	}
}

func show(api *spotify.API, name string) error {
	playlist, err := internal.SearchPlaylist(api, name)
	if err != nil {
		return err
	}

	if err := playlist.HREF.Get(api, playlist); err != nil {
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
