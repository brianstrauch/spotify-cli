package p

import (
	"errors"
	"spotify/internal"
	"spotify/internal/pause"
	"spotify/internal/play"
	"strings"

	"github.com/brianstrauch/spotify"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "p [song]",
		// Keep hidden, since this command is an alias.
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			var status string

			if len(args) > 0 {
				query := strings.Join(args, " ")
				status, err = play.Play(api, query)
			} else {
				status, err = p(api)
				if err != nil {
					return err
				}
			}

			cmd.Print(status)
			return nil
		},
	}
}

func p(api spotify.APIInterface) (string, error) {
	playback, err := api.Status()
	if err != nil {
		return "", err
	}

	if playback == nil {
		return "", errors.New(internal.NoActiveDeviceErr)
	}

	if playback.IsPlaying {
		return pause.Pause(api)
	} else {
		return play.Play(api, "")
	}
}
