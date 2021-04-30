package update

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"spotify/internal"
	"strings"

	"github.com/inconshreveable/go-update"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update to the latest version.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			version, err := updateFromGitHub("brianstrauch/spotify-cli")
			if err != nil {
				return err
			}

			cmd.Printf("Updated CLI to %s!\n", version)
			return nil
		},
	}
}

type Asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

type Release struct {
	TagName string  `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

func updateFromGitHub(repo string) (string, error) {
	release, err := getLatestReleaseInfo(repo)
	if err != nil {
		return "", err
	}

	// TODO: Prevent installing same release

	for _, asset := range release.Assets {
		if strings.Contains(asset.Name, runtime.GOOS) && strings.Contains(asset.Name, runtime.GOARCH) {
			res, err := http.Get(asset.BrowserDownloadURL)
			if err != nil {
				return "", err
			}
			defer res.Body.Close()

			// TODO: Checksum
			if err := update.Apply(res.Body, update.Options{}); err != nil {
				return "", err
			}

			return release.TagName, nil
		}
	}

	return "", errors.New(internal.NoReleaseAvailable)
}

func getLatestReleaseInfo(repo string) (*Release, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != err {
		return nil, err
	}

	release := new(Release)
	if err := json.Unmarshal(data, release); err != nil {
		return nil, err
	}

	return release, nil
}
