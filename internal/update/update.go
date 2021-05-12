package update

import (
	"spotify/pkg"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update to the latest version.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			version, err := pkg.UpdateFromGitHub("brianstrauch/spotify-cli")
			if err != nil {
				return err
			}

			cmd.Printf("Updated CLI to %s!\n", version)
			return nil
		},
	}
}
