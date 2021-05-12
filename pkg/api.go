package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"spotify/pkg/model"
	"strconv"
)

type APIInterface interface {
	Back() error
	Next() error
	Pause() error
	Play() error
	Save(id string) error
	Shuffle(state bool) error
	Status() (*model.Playback, error)
	Unsave(id string) error
}

type API struct {
	token string
}

func NewAPI(token string) *API {
	return &API{token}
}

func (a *API) Back() error {
	_, err := a.call(http.MethodPost, "/me/player/previous")
	return err
}

func (a *API) Next() error {
	_, err := a.call(http.MethodPost, "/me/player/next")
	return err
}

func (a *API) Pause() error {
	_, err := a.call(http.MethodPut, "/me/player/pause")
	return err
}

func (a *API) Play() error {
	_, err := a.call(http.MethodPut, "/me/player/play")
	return err
}

func (a *API) Save(id string) error {
	q := url.Values{}
	q.Add("ids", id)

	_, err := a.call(http.MethodPut, "/me/tracks?"+q.Encode())
	return err
}

func (a *API) Shuffle(state bool) error {
	q := url.Values{}
	q.Add("state", strconv.FormatBool(state))

	_, err := a.call(http.MethodPut, "/me/player/shuffle?"+q.Encode())
	return err
}

func (a *API) Status() (*model.Playback, error) {
	q := url.Values{}
	q.Add("additional_types", "episode")

	res, err := a.call(http.MethodGet, "/me/player?"+q.Encode())
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

func (a *API) Unsave(id string) error {
	q := url.Values{}
	q.Add("ids", id)

	_, err := a.call(http.MethodDelete, "/me/tracks?"+q.Encode())
	return err
}

func (a *API) call(method string, endpoint string) (*http.Response, error) {
	url := "https://api.spotify.com/v1" + endpoint

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.token))

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
