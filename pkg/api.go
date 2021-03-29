package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"spotify/pkg/model"
)

type SpotifyAPI struct {
	token string
}

func NewSpotifyAPI(token string) *SpotifyAPI {
	return &SpotifyAPI{token}
}

func (s *SpotifyAPI) Play() error {
	return s.call("PUT", "/me/player/play")
}

func (s *SpotifyAPI) Pause() error {
	return s.call("PUT", "/me/player/pause")
}

func (s *SpotifyAPI) call(method string, endpoint string) error {
	url := "https://api.spotify.com/v1" + endpoint

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token))

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
