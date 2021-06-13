package shuffle

import (
	"errors"
	"spotify/internal"

	"github.com/brianstrauch/spotify"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "shuffle",
		Short: "Toggle shuffle on or off.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			state, err := Shuffle(api)
			if err != nil {
				return err
			}

			if state {
				cmd.Println("🔀 Shuffle on")
			} else {
				cmd.Println("🔀 Shuffle off")
			}

			return nil
		},
	}
}

func Shuffle(api internal.APIInterface) (bool, error) {
	playback, err := api.GetPlayback()
	if err != nil {
		return false, nil
	}

	if playback == nil {
		return false, errors.New(internal.NoActiveDeviceErr)
	}

	state := playback.ShuffleState
	if err := api.Shuffle(!state); err != nil {
		return false, err
	}

	playback, err = internal.WaitForUpdatedPlayback(api, func(playback *spotify.Playback) bool {
		return playback.ShuffleState != state
	})
	if err != nil {
		return false, err
	}

	return playback.ShuffleState, nil
}
