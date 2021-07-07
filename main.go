package main

import (
	"spotify/internal/back"
	"spotify/internal/completion"
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
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
		Version:           "1.9.1",
		PersistentPreRunE: promptUpdate,
	}

	root.AddCommand(back.NewCommand())
	root.AddCommand(completion.NewCommand())
	root.AddCommand(login.NewCommand())
	root.AddCommand(next.NewCommand())
	root.AddCommand(p.NewCommand())
	root.AddCommand(pause.NewCommand())
	root.AddCommand(play.NewCommand())
	root.AddCommand(queue.NewCommand())
	root.AddCommand(repeat.NewCommand())
	root.AddCommand(save.NewCommand())
	root.AddCommand(shuffle.NewCommand())
	root.AddCommand(status.NewCommand())
	root.AddCommand(unsave.NewCommand())
	root.AddCommand(update.NewCommand())
	root.AddCommand(playlist.NewCommand())

	// Hide help command
	root.SetHelpCommand(&cobra.Command{Hidden: true})

	// Rename default flag descriptions
	root.Flags().BoolP("help", "h", false, "Help for Spotify CLI.")
	root.Flags().BoolP("version", "v", false, "Version for Spotify CLI.")

	err := root.Execute()
	cobra.CheckErr(err)
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
