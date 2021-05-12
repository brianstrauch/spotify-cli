package repeat

import (
	"errors"
	"spotify/internal"
	"spotify/pkg"
	"time"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "repeat",
		Short: "Toggle repeat on, off, or track.",
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
			case "off":
				cmd.Println("🔁 Repeat off")
			case "context":
				cmd.Println("🔁 Repeat on")
			case "track":
				cmd.Println("🔂 Repeat track")
			}

			return nil
		},
	}
}

func Repeat(api pkg.APIInterface) (string, error) {
	playback, err := api.Status()
	if err != nil {
		return "", nil
	}

	if playback == nil {
		return "", errors.New(internal.NoActiveDeviceErr)
	}

	state := playback.RepeatState
	if err := api.Repeat(toggle(state)); err != nil {
		return "", err
	}

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

			if playback.RepeatState != state {
				return playback.RepeatState, nil
			}
		}
	}
}

func toggle(state string) string {
	switch state {
	case "off":
		return "context"
	case "context":
		return "track"
	case "track":
		return "off"
	}

	return ""
}
