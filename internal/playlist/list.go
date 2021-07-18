package playlist

import (
	"errors"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"spotify/internal"
	"strconv"
	"strings"

	"github.com/brianstrauch/spotify"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list your playlists",
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			list, err := List(api)
			if err != nil {
				return err
			}

			fmt.Print(list)
			return nil
		},
	}
}

func List(api *spotify.API) (string, error) {
	playlists, err := api.GetPlaylists()
	if err != nil {
		return "", err
	}

	if len(playlists) == 0 {
		return "", errors.New(internal.ErrNoPlaylists)
	}

	output := new(strings.Builder)

	table := tablewriter.NewWriter(output)
	table.SetBorder(false)

	table.SetHeader([]string{"#", "Playlist"})

	for i, playlist := range playlists {
		table.Append([]string{strconv.Itoa(i + 1), playlist.Name})
	}
	table.Render()

	return output.String(), nil
}
