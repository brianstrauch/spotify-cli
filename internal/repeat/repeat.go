package repeat

import (
	"errors"
	"spotify/internal"

	"github.com/brianstrauch/spotify"
	"github.com/brianstrauch/spotify/model"
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
		Short: "Cycle repeat through on, off, or track.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			state, err := Repeat(api)
			if err != nil {
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

func Repeat(api spotify.APIInterface) (string, error) {
	playback, err := api.Status()
	if err != nil {
		return "", nil
	}

	if playback == nil {
		return "", errors.New(internal.NoActiveDeviceErr)
	}

	state := playback.RepeatState
	if err := api.Repeat(cycle(state)); err != nil {
		return "", err
	}

	playback, err = api.WaitForUpdatedPlayback(func(playback *model.Playback) bool {
		return playback.RepeatState != state
	})
	if err != nil {
		return "", err
	}

	return playback.RepeatState, nil
}

func cycle(state string) string {
	for i := range states {
		if states[i] == state {
			j := (i + 1) % len(states)
			return states[j]
		}
	}
	return ""
}
