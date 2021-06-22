package repeat

import (
	"errors"
	"spotify/internal"

	"github.com/brianstrauch/spotify"
	"github.com/spf13/cobra"
)

const (
	StateOff   = "off"
	StateOn    = "context"
	StateTrack = "track"
)

var states = []string{StateOff, StateOn, StateTrack}

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "repeat",
		Short: "Set repeat to on, off, or track.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			var state string

			switch args[0] {
			case "on":
				state = StateOn
			case "off":
				state = StateOff
			case "track":
				state = StateTrack
			default:
				return errors.New(internal.RepeatArgErr)
			}

			if err := Repeat(api, state); err != nil {
				return err
			}

			switch state {
			case StateOff:
				cmd.Println("🔁 Repeat off")
			case StateOn:
				cmd.Println("🔁 Repeat on")
			case StateTrack:
				cmd.Println("🔂 Repeat track")
			}

			return nil
		},
	}
}

func Repeat(api internal.APIInterface, state string) error {
	if err := api.Repeat(state); err != nil {
		return err
	}

	_, err := internal.WaitForUpdatedPlayback(api, func(playback *spotify.Playback) bool {
		return playback.RepeatState == state
	})
	return err
}
