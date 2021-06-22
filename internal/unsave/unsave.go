package unsave

import (
	"errors"
	"spotify/internal"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "unsave",
		Short: "Unsave the current song.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			if err := unsave(api); err != nil {
				return err
			}

			cmd.Println("Unsaved.")
			return nil
		},
	}
}

func unsave(api internal.APIInterface) error {
	playback, err := api.GetPlayback()
	if err != nil {
		return err
	}

	if playback.Item.Type == "episode" {
		return errors.New(internal.ErrSavePodcast)
	}

	return api.RemoveSavedTracks(playback.Item.ID)
}
