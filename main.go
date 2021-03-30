package main

import (
	"fmt"
	"spotify/internal"

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
		Use:              CommandName,
		Short:            "Play music from the command line.",
		PersistentPreRun: update,
	}

	root.AddCommand(internal.NewLoginCommand())
	root.AddCommand(internal.NewPlayCommand())
	root.AddCommand(internal.NewPauseCommand())
	root.AddCommand(internal.NewVersionCommand())

	// Hide help command and rename help flag
	root.SetHelpCommand(&cobra.Command{Hidden: true})
	root.Flags().BoolP("help", "h", false, fmt.Sprintf("Help for %s.", FullName))

	root.Execute()
}

func update(cmd *cobra.Command, _ []string) {
	// TODO: Check for updates
}
