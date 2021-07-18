package shuffle

import (
	"spotify/internal"

	"github.com/brianstrauch/spotify"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:       "shuffle",
		Short:     "turn shuffle on or off",
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{"on", "off"},
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			var state bool

			switch args[0] {
			case "on":
				state = true
			case "off":
				state = false
			}

			if err := Shuffle(api, state); err != nil {
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

func Shuffle(api internal.APIInterface, state bool) error {
	if err := api.Shuffle(state); err != nil {
		return err
	}

	_, err := internal.WaitForUpdatedPlayback(api, func(playback *spotify.Playback) bool {
		return playback.ShuffleState == state
	})
	return err
}
