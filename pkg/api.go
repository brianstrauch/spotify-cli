package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"spotify/pkg/model"
	"strconv"
	"time"
)

type APIInterface interface {
	Back() error
	Next() error
	Pause() error
	Play(uri string) error
	Repeat(state string) error
	Save(id string) error
	Search(queue string, limit int) (*model.Page, error)
	Shuffle(state bool) error
	Status() (*model.Playback, error)
	Unsave(id string) error
	WaitForUpdatedPlayback(isUpdated func(*model.Playback) bool) (*model.Playback, error)
}

type API struct {
	token string
}

func NewAPI(token string) *API {
	return &API{token}
}

func (a *API) Back() error {
	_, err := a.call(http.MethodPost, "/me/player/previous", nil)
	return err
}

func (a *API) Next() error {
	_, err := a.call(http.MethodPost, "/me/player/next", nil)
	return err
}

func (a *API) Pause() error {
	_, err := a.call(http.MethodPut, "/me/player/pause", nil)
	return err
}

func (a *API) Play(uri string) error {
	if len(uri) == 0 {
		_, err := a.call(http.MethodPut, "/me/player/play", nil)
		return err
	}

	type Body struct {
		URIs []string `json:"uris"`
	}

	body := new(Body)
	body.URIs = []string{uri}

	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = a.call(http.MethodPut, "/me/player/play", bytes.NewReader(data))
	return err
}

func (a *API) Repeat(state string) error {
	q := url.Values{}
	q.Add("state", state)

	_, err := a.call(http.MethodPut, "/me/player/repeat?"+q.Encode(), nil)
	return err
}

func (a *API) Save(id string) error {
	q := url.Values{}
	q.Add("ids", id)

	_, err := a.call(http.MethodPut, "/me/tracks?"+q.Encode(), nil)
	return err
}

func (a *API) Search(query string, limit int) (*model.Page, error) {
	q := url.Values{}
	q.Add("q", query)
	q.Add("type", "track")
	q.Add("limit", strconv.Itoa(limit))

	res, err := a.call(http.MethodGet, "/search?"+q.Encode(), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	page := new(model.Page)
	err = json.NewDecoder(res.Body).Decode(page)

	return page, err
}

func (a *API) Shuffle(state bool) error {
	q := url.Values{}
	q.Add("state", strconv.FormatBool(state))

	_, err := a.call(http.MethodPut, "/me/player/shuffle?"+q.Encode(), nil)
	return err
}

func (a *API) Status() (*model.Playback, error) {
	q := url.Values{}
	q.Add("additional_types", "episode")

	res, err := a.call(http.MethodGet, "/me/player?"+q.Encode(), nil)
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

	_, err := a.call(http.MethodDelete, "/me/tracks?"+q.Encode(), nil)
	return err
}

func (a *API) WaitForUpdatedPlayback(isUpdated func(playback *model.Playback) bool) (*model.Playback, error) {
	timeout := time.After(time.Second)
	tick := time.Tick(100 * time.Millisecond)

	for {
		select {
		case <-timeout:
			return nil, errors.New("request timed out")
		case <-tick:
			playback, err := a.Status()
			if err != nil {
				return nil, err
			}

			if isUpdated(playback) {
				return playback, nil
			}
		}
	}
}

func (a *API) call(method string, endpoint string, body io.Reader) (*http.Response, error) {
	url := "https://api.spotify.com/v1" + endpoint

	req, err := http.NewRequest(method, url, body)
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
