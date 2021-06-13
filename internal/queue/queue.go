package queue

import (
	"spotify/internal"
	"strings"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "queue song",
		Aliases: []string{"q"},
		Short:   "Queue a specific song.",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			query := strings.Join(args, " ")

			if err := Queue(api, query); err != nil {
				return err
			}

			cmd.Println("Queued!")
			return nil
		},
	}
}

func Queue(api internal.APIInterface, query string) error {
	uri, err := internal.Search(api, query)
	if err != nil {
		return err
	}

	return api.Queue(uri)
}
