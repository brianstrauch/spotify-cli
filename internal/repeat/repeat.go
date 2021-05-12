package repeat

import (
	"errors"
	"spotify/internal"
	"spotify/pkg"
	"spotify/pkg/model"

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

	playback, err = api.WaitForUpdatedPlayback(func(playback *model.Playback) bool {
		return playback.RepeatState != state
	})
	if err != nil {
		return "", err
	}

	return playback.RepeatState, nil
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
