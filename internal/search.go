package internal

import "spotify/pkg"

func Search(api pkg.APIInterface, query string) (string, error) {
	page, err := api.Search(query, 1)
	if err != nil {
		return "", err
	}

	return page.Tracks.Items[0].URI, nil
}
