package status

import (
	"errors"
	"spotify/internal"
	"spotify/pkg"
	"spotify/pkg/model"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "status",
		Aliases: []string{"s"},
		Short:   "Show the current song.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			status, err := status(api)
			if err != nil {
				return err
			}

			cmd.Print(status)
			return nil
		},
	}
}

func Show(playback *model.Playback) string {
	artists := playback.Item.Artists

	status := playback.Item.Name + "\n"
	status += artists[0].Name
	for i := 1; i < len(artists); i++ {
		status += ", " + artists[i].Name
	}
	status += "\n"

	return status
}

func status(api pkg.APIInterface) (string, error) {
	playback, err := api.Status()
	if err != nil {
		return "", err
	}

	if playback == nil {
		return "", errors.New(internal.NoActiveDeviceErr)
	}

	return Show(playback), nil
}
