package unsave

import (
	"errors"
	"spotify/internal"
	"spotify/pkg"

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

func unsave(api pkg.APIInterface) error {
	playback, err := api.Status()
	if err != nil {
		return err
	}

	if playback.Item.Type == "episode" {
		return errors.New(internal.SavePodcastErr)
	}

	return api.Unsave(playback.Item.ID)
}
