package device

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "device",
		Short: "Manage devices.",
	}

	cmd.AddCommand(NewListCommand())

	return cmd
}
