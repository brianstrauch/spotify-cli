package playlist

import (
	"fmt"
	"spotify/internal"

	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "List playlists.",
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}
			playlists, err := api.GetPlaylists()
			if err != nil {
				return err
			}
			for _, pl := range playlists {
				fmt.Println(pl.Name)
			}
			return nil
		},
	}
}
