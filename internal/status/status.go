package status

import (
	"spotify/internal"
	"spotify/pkg"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show the current song.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			status, err := status(api)
			if err != nil {
				return err
			}

			cmd.Println(status)
			return nil
		},
	}
}

func status(api pkg.APIInterface) (string, error) {
	playback, err := api.Status()
	if err != nil {
		return "", err
	}

	artists := playback.Item.Artists

	status := playback.Item.Name + "\n"
	status += artists[0].Name
	for i := 1; i < len(artists); i++ {
		status += ", " + artists[i].Name
	}

	return status, nil
}
