package next

import (
	"errors"
	"spotify/internal"
	"spotify/internal/status"
	"spotify/pkg"
	"time"

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

	timeout := time.After(time.Second)
	tick := time.Tick(100 * time.Millisecond)

	for {
		select {
		case <-timeout:
			return "", nil
		case <-tick:
			playback, err := api.Status()
			if err != nil {
				return "", err
			}

			if id != playback.Item.ID {
				return status.Show(playback), nil
			}
		}
	}
}
