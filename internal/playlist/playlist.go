package playlist

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "playlist",
		Short: "Manage playlists.",
	}

	cmd.AddCommand(NewListCommand())
	cmd.AddCommand(NewShowCommand())

	return cmd
}
