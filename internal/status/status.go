package status

import (
	"errors"
	"fmt"
	"spotify/internal"
	"spotify/pkg"
	"spotify/pkg/model"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "status",
		Aliases: []string{"s"},
		Short:   "Show the current song or episode.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			status, err := status(api)
			if err != nil {
				return err
			}

			cmd.Print(status)
			return nil
		},
	}
}

func status(api pkg.APIInterface) (string, error) {
	playback, err := api.Status()
	if err != nil {
		return "", err
	}

	if playback == nil {
		return "", errors.New(internal.NoActiveDeviceErr)
	}

	return Show(playback), nil
}

func Show(playback *model.Playback) string {
	status := fmt.Sprintf("🎵 %s\n", playback.Item.Name)

	switch playback.Item.Type {
	case "track":
		status += fmt.Sprintf("🎤 %s\n", joinArtists(playback.Item.Artists))
	case "episode":
		status += fmt.Sprintf("🎤 %s\n", playback.Item.Show.Name)
	}

	if playback.IsPlaying {
		status += "▶️  "
	} else {
		status += "⏸  "
	}

	status += showProgressBar(playback.ProgressMs, playback.Item.DurationMs)

	return status
}

func joinArtists(artists []model.Artist) string {
	list := artists[0].Name
	for i := 1; i < len(artists); i++ {
		list += ", " + artists[i].Name
	}
	return list
}

func showProgressBar(progress, duration int) string {
	const length = 16
	bars := length * progress / duration

	status := fmt.Sprintf("%s [", formatTime(progress))
	for i := 0; i < bars; i++ {
		status += "="
	}
	for i := bars; i < length; i++ {
		status += " "
	}
	status += fmt.Sprintf("] %s\n", formatTime(duration))

	return status
}

func formatTime(ms int) string {
	s := ms / 1000
	return fmt.Sprintf("%d:%02d", s/60, s%60)
}
