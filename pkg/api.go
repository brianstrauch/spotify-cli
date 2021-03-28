package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"spotify/pkg/model"
)

const BaseAPIURL = "https://api.spotify.com/v1"

func Play(token string) error {
	return call("PUT", "/me/player/play", token)
}

func Pause(token string) error {
	return call("PUT", "/me/player/pause", token)
}

func call(method string, endpoint string, token string) error {
	req, err := http.NewRequest(method, BaseAPIURL+endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// Success
	if res.StatusCode == http.StatusNoContent {
		return nil
	}

	// Error
	spotifyErr := new(model.SpotifyError)
	if err := json.NewDecoder(res.Body).Decode(spotifyErr); err != nil {
		return err
	}

	return errors.New(spotifyErr.Error.Message)
}
