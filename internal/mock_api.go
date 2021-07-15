package internal

import (
	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/mock"
)

type APIInterface interface {
	SaveTracks(ids ...string) error
	RemoveSavedTracks(ids ...string) error

	GetPlayback() (*spotify.Playback, error)
	Play(deviceID string, uris ...string) error
	Pause(deviceID string) error
	SkipToNextTrack() error
	SkipToPreviousTrack() error
	Repeat(state string) error
	Shuffle(state bool) error
	Queue(uri string) error

	Search(q string, limit int) (*spotify.Paging, error)
}

type MockAPI struct {
	mock.Mock
}

func (m *MockAPI) SaveTracks(ids ...string) error {
	args := m.Called(ids)
	return args.Error(0)
}

func (m *MockAPI) RemoveSavedTracks(ids ...string) error {
	args := m.Called(ids)
	return args.Error(0)
}

func (m *MockAPI) GetPlayback() (*spotify.Playback, error) {
	args := m.Called()

	playback := args.Get(0)
	err := args.Error(1)

	if playback == nil {
		return nil, err
	}

	return playback.(*spotify.Playback), err
}

func (m *MockAPI) Play(deviceID string, uris ...string) error {
	args := m.Called(deviceID, uris)
	return args.Error(0)
}

func (m *MockAPI) Pause(deviceID string) error {
	args := m.Called(deviceID)
	return args.Error(0)
}

func (m *MockAPI) SkipToNextTrack() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockAPI) SkipToPreviousTrack() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockAPI) Repeat(state string) error {
	args := m.Called(state)
	return args.Error(0)
}

func (m *MockAPI) Shuffle(state bool) error {
	args := m.Called(state)
	return args.Error(0)
}

func (m *MockAPI) Queue(uri string) error {
	args := m.Called(uri)
	return args.Error(0)
}

func (m *MockAPI) Search(q string, limit int) (*spotify.Paging, error) {
	args := m.Called(q, limit)

	page := args.Get(0)
	err := args.Error(1)

	if page == nil {
		return nil, err
	}

	return page.(*spotify.Paging), err
}
