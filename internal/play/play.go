package play

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
		Use:   "play",
		Short: "Play music.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			status, err := Play(api)
			if err != nil {
				return err
			}

			cmd.Print(status)
			return nil
		},
	}
}

func Play(api pkg.APIInterface) (string, error) {
	playback, err := api.Status()
	if err != nil {
		return "", err
	}

	if playback == nil {
		return "", errors.New(internal.NoActiveDeviceErr)
	}

	if err := api.Play(); err != nil {
		if err.Error() == internal.RestrictionViolatedSpotifyErr {
			return "", errors.New(internal.AlreadyPlayingErr)
		}
	}

	for {
		select {
		case <-time.After(time.Second):
			return "", nil
		case <-time.Tick(100 * time.Millisecond):
			playback, err := api.Status()
			if err != nil {
				return "", err
			}

			if playback.IsPlaying {
				return status.Show(playback), nil
			}
		}
	}
}
