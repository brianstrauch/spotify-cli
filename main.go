package main

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"spotify/internal/back"
	"spotify/internal/login"
	"spotify/internal/next"
	"spotify/internal/p"
	"spotify/internal/pause"
	"spotify/internal/play"
	"spotify/internal/playlist"
	"spotify/internal/queue"
	"spotify/internal/repeat"
	"spotify/internal/save"
	"spotify/internal/shuffle"
	"spotify/internal/status"
	"spotify/internal/unsave"
	"spotify/internal/update"
)

// version is a linker flag set by goreleaser
var version = "0.0.0"

func main() {
	// TODO: https://github.com/spf13/viper/pull/1064
	viper.AddConfigPath("$HOME")
	viper.SetConfigName(".spotify-cli")
	viper.SetConfigType("json")
	_ = viper.SafeWriteConfig()
	_ = viper.ReadInConfig()

	root := &cobra.Command{
		Use:               "spotify",
		Short:             "Spotify for the terminal 🎵",
		Version:           version,
		PersistentPreRunE: promptUpdate,
	}

	root.AddCommand(back.NewCommand())
	root.AddCommand(login.NewCommand())
	root.AddCommand(next.NewCommand())
	root.AddCommand(p.NewCommand())
	root.AddCommand(pause.NewCommand())
	root.AddCommand(play.NewCommand())
	root.AddCommand(playlist.NewCommand())
	root.AddCommand(queue.NewCommand())
	root.AddCommand(repeat.NewCommand())
	root.AddCommand(save.NewCommand())
	root.AddCommand(shuffle.NewCommand())
	root.AddCommand(status.NewCommand())
	root.AddCommand(unsave.NewCommand())
	root.AddCommand(update.NewCommand())

	root.SetHelpCommand(&cobra.Command{Hidden: true})
	root.SetVersionTemplate("v{{.Version}}\n")

	_ = root.Execute()
}

func promptUpdate(cmd *cobra.Command, _ []string) error {
	if time.Now().Unix() < viper.GetInt64("prompt_update_timer") {
		return nil
	}

	isUpdated, err := update.IsUpdated(cmd)
	if err != nil {
		return err
	}
	if !isUpdated {
		cmd.Println("Use 'spotify update' to get the latest version.")
	}

	// Wait one day before the next prompt
	const day int64 = 24 * 60 * 60
	viper.Set("prompt_update_timer", time.Now().Unix()+day)

	return viper.WriteConfig()
}
