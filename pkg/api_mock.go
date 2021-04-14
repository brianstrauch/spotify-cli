package pkg

import (
	"spotify/pkg/model"

	"github.com/stretchr/testify/mock"
)

type MockSpotifyAPI struct {
	mock.Mock
}

func (m *MockSpotifyAPI) Back() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockSpotifyAPI) Next() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockSpotifyAPI) Pause() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockSpotifyAPI) Play() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockSpotifyAPI) Status() (*model.Playback, error) {
	args := m.Called()

	playback := args.Get(0)
	err := args.Error(1)

	if playback == nil {
		return nil, err
	}

	return playback.(*model.Playback), err

}
