package pkg

import (
	"spotify/pkg/model"

	"github.com/stretchr/testify/mock"
)

type MockSpotifyAPI struct {
	mock.Mock
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

func (m *MockSpotifyAPI) Previous() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockSpotifyAPI) Status() (*model.Playback, error) {
	args := m.Called()
	return args.Get(0).(*model.Playback), args.Error(1)
}
