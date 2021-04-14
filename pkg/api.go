package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"spotify/pkg/model"
)

type APIInterface interface {
	Back() error
	Next() error
	Pause() error
	Play() error
	Status() (*model.Playback, error)
}

type API struct {
	token string
}

func NewAPI(token string) *API {
	return &API{token}
}

func (s *API) Back() error {
	_, err := s.call("POST", "/me/player/previous")
	return err
}

func (s *API) Next() error {
	_, err := s.call("POST", "/me/player/next")
	return err
}

func (s *API) Pause() error {
	_, err := s.call("PUT", "/me/player/pause")
	return err
}

func (s *API) Play() error {
	_, err := s.call("PUT", "/me/player/play")
	return err
}

func (s *API) Status() (*model.Playback, error) {
	res, err := s.call("GET", "/me/player")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	playback := new(model.Playback)
	err = json.NewDecoder(res.Body).Decode(playback)

	return playback, err
}

func (s *API) call(method string, endpoint string) (*http.Response, error) {
	url := "https://api.spotify.com/v1" + endpoint

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token))

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Success
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return res, nil
	}

	// Error
	spotifyErr := new(model.SpotifyError)
	if err := json.NewDecoder(res.Body).Decode(spotifyErr); err != nil {
		return nil, err
	}

	return nil, errors.New(spotifyErr.Error.Message)
}
