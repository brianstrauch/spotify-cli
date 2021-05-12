package pkg

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"

	"github.com/inconshreveable/go-update"
)

type Asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

type Release struct {
	TagName string  `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

func UpdateFromGitHub(repo string) (string, error) {
	release, err := getLatestReleaseInfo(repo)
	if err != nil {
		return "", err
	}

	checksums, err := getChecksumsFromRelease(release)
	if err != nil {
		return "", err
	}

	for _, asset := range release.Assets {
		if strings.Contains(asset.Name, runtime.GOOS) && strings.Contains(asset.Name, runtime.GOARCH) {
			res, err := http.Get(asset.BrowserDownloadURL)
			if err != nil {
				return "", err
			}
			defer res.Body.Close()

			options := update.Options{Checksum: checksums[asset.Name]}
			if err := update.Apply(res.Body, options); err != nil {
				return "", err
			}

			return release.TagName, nil
		}
	}

	return "", errors.New("No release available for this OS and architecture")
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

func getChecksumsFromRelease(release *Release) (map[string][]byte, error) {
	checksums := make(map[string][]byte)

	for _, asset := range release.Assets {
		if strings.Contains(asset.Name, "checksums.txt") {
			res, err := http.Get(asset.BrowserDownloadURL)
			if err != nil {
				return nil, err
			}
			defer res.Body.Close()

			out, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}
			lines := strings.TrimRight(string(out), "\n")

			for _, line := range strings.Split(lines, "\n") {
				x := strings.Split(line, "  ")

				checksum, err := hex.DecodeString(x[0])
				if err != nil {
					return nil, err
				}

				file := x[1]
				checksums[file] = checksum
			}
		}
	}

	return checksums, nil
}
