package playlist

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "playlist",
		Aliases: []string{"p"},
		Short:   "do things with playlists",
	}
	cmd.AddCommand(NewListCommand())
	cmd.AddCommand(NewDetailsCommand())
	return cmd
}
