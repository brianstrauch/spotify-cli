package device

import (
	"errors"
	"fmt"
	"spotify/internal"
	"strings"

	"github.com/brianstrauch/spotify"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list your devices",
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			list, err := List(api)
			if err != nil {
				return err
			}

			fmt.Print(list)
			return nil
		},
	}
}

func List(api *spotify.API) (string, error) {
	devices, err := api.GetDevices()
	if err != nil {
		return "", err
	}

	if len(devices) == 0 {
		return "", errors.New(internal.ErrNoDevices)
	}

	output := new(strings.Builder)

	table := tablewriter.NewWriter(output)
	table.SetBorder(false)

	table.SetHeader([]string{"ID", "Device"})

	for _, device := range devices {
		table.Append([]string{device.ID, device.Name})
	}
	table.Render()

	return output.String(), nil
}
