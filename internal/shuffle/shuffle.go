package shuffle

import (
	"errors"
	"spotify/internal"
	"spotify/pkg"
	"time"

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

func Shuffle(api pkg.APIInterface) (bool, error) {
	playback, err := api.Status()
	if err != nil {
		return false, nil
	}

	if playback == nil {
		return false, errors.New(internal.NoActiveDeviceErr)
	}

	isShuffled := playback.ShuffleState
	if err := api.Shuffle(!isShuffled); err != nil {
		return false, err
	}

	timeout := time.After(time.Second)
	tick := time.Tick(100 * time.Millisecond)

	for {
		select {
		case <-timeout:
			return false, nil
		case <-tick:
			playback, err := api.Status()
			if err != nil {
				return false, err
			}

			if playback.ShuffleState != isShuffled {
				return playback.ShuffleState, nil
			}
		}
	}
}
