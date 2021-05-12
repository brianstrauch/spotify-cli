package next

import (
	"errors"
	"spotify/internal"
	"spotify/internal/status"
	"spotify/pkg"
	"spotify/pkg/model"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "next",
		Aliases: []string{"n"},
		Short:   "Skip to next song.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			status, err := next(api)
			if err != nil {
				return err
			}

			cmd.Print(status)
			return nil
		},
	}
}

func next(api pkg.APIInterface) (string, error) {
	playback, err := api.Status()
	if err != nil {
		return "", err
	}

	if playback == nil {
		return "", errors.New(internal.NoActiveDeviceErr)
	}

	id := playback.Item.ID

	if err := api.Next(); err != nil {
		return "", err
	}

	playback, err = api.WaitForUpdatedPlayback(func(playback *model.Playback) bool {
		return playback.Item.ID != id
	})
	if err != nil {
		return "", err
	}

	return status.Show(playback), nil
}
