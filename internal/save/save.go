package save

import (
	"errors"
	"spotify/internal"

	"github.com/brianstrauch/spotify"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "save",
		Short: "Save the current song.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			if err := save(api); err != nil {
				return err
			}

			cmd.Println("Saved!")
			return nil
		},
	}
}

func save(api spotify.APIInterface) error {
	playback, err := api.GetPlayback()
	if err != nil {
		return err
	}

	if playback.Item.Type == "episode" {
		return errors.New(internal.SavePodcastErr)
	}

	return api.SaveTracks(playback.Item.ID)
}
