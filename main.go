package main

import (
	"log"
	"spotify/internal/back"
	"spotify/internal/login"
	"spotify/internal/next"
	"spotify/internal/p"
	"spotify/internal/pause"
	"spotify/internal/play"
	"spotify/internal/queue"
	"spotify/internal/repeat"
	"spotify/internal/save"
	"spotify/internal/shuffle"
	"spotify/internal/status"
	"spotify/internal/unsave"
	"spotify/internal/update"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	// TODO: https://github.com/spf13/viper/pull/1064
	viper.AddConfigPath("$HOME")
	viper.SetConfigName(".spotify-cli")
	viper.SetConfigType("json")
	viper.SafeWriteConfig()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	root := &cobra.Command{
		Use:     "spotify",
		Short:   "Spotify for the terminal 🎵",
		Version: "1.5.1",
	}

	root.AddCommand(back.NewCommand())
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

	// Hide help command
	root.SetHelpCommand(&cobra.Command{Hidden: true})

	// Rename default flag descriptions
	root.Flags().BoolP("help", "h", false, "Help for Spotify CLI.")
	root.Flags().BoolP("version", "v", false, "Version for Spotify CLI.")

	root.Execute()
}
