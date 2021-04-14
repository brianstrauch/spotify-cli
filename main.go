package main

import (
	"fmt"
	"spotify/internal/back"
	"spotify/internal/login"
	"spotify/internal/next"
	"spotify/internal/p"
	"spotify/internal/pause"
	"spotify/internal/play"
	"spotify/internal/status"
	"spotify/internal/version"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	CommandName = "spotify"
	FullName    = "Spotify CLI"
)

func main() {
	// TODO: https://github.com/spf13/viper/pull/1064
	viper.AddConfigPath("$HOME")
	viper.SetConfigName(".spotify-cli")
	viper.SetConfigType("json")
	viper.SafeWriteConfig()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	root := &cobra.Command{
		Use:   CommandName,
		Short: "Play music from the command line.",
	}

	root.AddCommand(back.NewCommand())
	root.AddCommand(login.NewCommand())
	root.AddCommand(next.NewCommand())
	root.AddCommand(p.NewCommand())
	root.AddCommand(pause.NewCommand())
	root.AddCommand(play.NewCommand())
	root.AddCommand(status.NewCommand())
	root.AddCommand(version.NewCommand())

	// Hide help command and rename help flag
	root.SetHelpCommand(&cobra.Command{Hidden: true})
	root.Flags().BoolP("help", "h", false, fmt.Sprintf("Help for %s.", FullName))

	root.Execute()
}
