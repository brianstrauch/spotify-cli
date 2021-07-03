package playlist

import (
	"errors"
	"fmt"
	"spotify/internal"

	"github.com/brianstrauch/spotify"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List your playlists.",
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

	list := ""
	for _, playlist := range playlists {
		list += playlist.Name + "\n"
	}

	return list, nil
}
