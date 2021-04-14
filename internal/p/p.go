package p

import (
	"errors"
	"spotify/internal"
	"spotify/internal/pause"
	"spotify/internal/play"
	"spotify/pkg"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "p",
		// Keep hidden, since this command is an alias.
		Hidden: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			return p(api)
		},
	}
}

func p(api pkg.APIInterface) error {
	playback, err := api.Status()
	if err != nil {
		return err
	}

	if playback == nil {
		return errors.New(internal.NoActiveDeviceErr)
	}

	if playback.IsPlaying {
		return pause.Pause(api)
	} else {
		return play.Play(api)
	}
}
