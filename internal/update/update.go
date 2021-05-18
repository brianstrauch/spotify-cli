package update

import (
	"errors"
	"spotify/internal"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
)

const repo = "brianstrauch/spotify-cli"

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update to the latest version.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			current, err := semver.Parse(cmd.Root().Version)
			if err != nil {
				return err
			}

			latest, found, err := selfupdate.DetectLatest(repo)
			if err != nil {
				return err
			}

			if !found || current.Equals(latest.Version) {
				return errors.New(internal.AlreadyUpToDateErr)
			}

			if _, err := selfupdate.UpdateSelf(current, repo); err != nil {
				return err
			}

			cmd.Printf("Updated CLI to version %s!\n", latest.Version.String())
			return nil
		},
	}
}
