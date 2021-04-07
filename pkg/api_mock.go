package pkg

import "spotify/pkg/model"

type MockSpotifyAPI struct{}

func (m *MockSpotifyAPI) Play() error {
	return nil
}

func (m *MockSpotifyAPI) Pause() error {
	return nil
}

func (m *MockSpotifyAPI) Status() (*model.Playback, error) {
	return nil, nil
}
