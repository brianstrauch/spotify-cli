package main

import (
	_ "embed"
	"fmt"
	"spotify/internal"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	CommandName = "spotify"
	FullName    = "Spotify CLI"
)

//go:embed version.txt
var version string

func main() {
	viper.SetConfigFile("config.json")
	viper.SafeWriteConfig()
	viper.ReadInConfig()

	root := &cobra.Command{
		Use:              CommandName,
		Short:            "Play music from the command line.",
		PersistentPreRun: update,
	}

	root.AddCommand(internal.NewLoginCommand())
	root.AddCommand(internal.NewPlayCommand())
	root.AddCommand(internal.NewPauseCommand())

	// Hide help command and rename help flag
	root.SetHelpCommand(&cobra.Command{Hidden: true})
	root.Flags().BoolP("help", "h", false, fmt.Sprintf("Help for %s.", FullName))

	root.Execute()
}

func update(cmd *cobra.Command, _ []string) {
	// TODO: Check for updates
}
