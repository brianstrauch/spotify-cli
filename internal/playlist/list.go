package playlist

import (
	"fmt"
	"spotify/internal"
	"strings"

	"github.com/brianstrauch/spotify"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List playlists.",
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}
			output, err := List(api)
			if err != nil {
				return err
			}
			fmt.Print(output)
			return nil
		},
	}
}

func List(api *spotify.API) (string, error) {
	playlists, err := api.GetPlaylists()
	if err != nil {
		return "", err
	}
	var builder strings.Builder
	for _, pl := range playlists {
		builder.WriteString(fmt.Sprintln(pl.Name))
	}
	return builder.String(), nil
}
