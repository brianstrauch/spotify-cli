package playlist

import (
	"errors"
	"fmt"
	"spotify/internal"
	"strings"

	"github.com/spf13/cobra"
)

func NewDetailsCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show playlist artist and songs.",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}
			playlists, err := api.GetPlaylists()
			if err != nil {
				return err
			}
			id := ""
			for _, playlist := range playlists {
				if strings.ToLower(playlist.Name) == strings.ToLower(args[0]) {
					id = playlist.ID
				}
			}
			if playlistID == "" {
				return errors.New("no such playlist")
			}
			playlist, err := api.GetPlaylist(playlistID)
			if err != nil {
				return err
			}
			fmt.Printf("📝: %s\n", playlist.Name)
			fmt.Println("💿 Tracks:")
			for i, tr := range playlist.Tracks.Items {
				artistNames := make([]string, len(tr.Track.Artists))
				for i, artist := range tr.Track.Artists {
					if err := artist.Get(api, &artist); err != nil {
						return err
					}
					artistNames[i] = artist.Name
				}
				fmt.Printf("%d. %s - %s\n", i+1, strings.Join(artistNames, ". "), tr.Track.Name)
			}
			return nil
		},
	}
}
