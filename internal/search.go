package internal

import "github.com/brianstrauch/spotify"

func Search(api spotify.APIInterface, query string) (string, error) {
	page, err := api.Search(query, 1)
	if err != nil {
		return "", err
	}

	return page.Tracks.Items[0].URI, nil
}
