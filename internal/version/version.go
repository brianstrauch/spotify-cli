package version

import "github.com/spf13/cobra"

const version = "1.0.0"

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show the version of Spotify CLI",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Println(version)
		},
	}
}
