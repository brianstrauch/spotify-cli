package internal

import (
	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/mock"
)

type APIInterface interface {
	SaveTracks(ids ...string) error
	RemoveSavedTracks(ids ...string) error

	GetPlayback() (*spotify.Playback, error)
	GetDevices() ([]*spotify.Device, error)
	Play(contextURI string, uris ...string) error
	Pause() error
	SkipToNextTrack() error
	SkipToPreviousTrack() error
	Repeat(state string) error
	Shuffle(state bool) error
	Queue(uri string) error

	GetPlaylists() ([]*spotify.Playlist, error)

	Search(q, searchType string, limit int) (*spotify.Paging, error)
}

type MockAPI struct {
	mock.Mock
}

func (m *MockAPI) SaveTracks(ids ...string) error {
	return m.Called(ids).Error(0)
}

func (m *MockAPI) RemoveSavedTracks(ids ...string) error {
	return m.Called(ids).Error(0)
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

func (m *MockAPI) GetDevices() ([]*spotify.Device, error) {
	args := m.Called()

	devices := args.Get(0)
	err := args.Error(1)

	return devices.([]*spotify.Device), err
}

func (m *MockAPI) Play(contextURI string, uris ...string) error {
	return m.Called(contextURI, uris).Error(0)
}

func (m *MockAPI) Pause() error {
	return m.Called().Error(0)
}

func (m *MockAPI) SkipToNextTrack() error {
	return m.Called().Error(0)
}

func (m *MockAPI) SkipToPreviousTrack() error {
	return m.Called().Error(0)
}

func (m *MockAPI) Repeat(state string) error {
	return m.Called(state).Error(0)
}

func (m *MockAPI) Shuffle(state bool) error {
	return m.Called(state).Error(0)
}

func (m *MockAPI) Queue(uri string) error {
	return m.Called(uri).Error(0)
}

func (m *MockAPI) GetPlaylists() ([]*spotify.Playlist, error) {
	args := m.Called()

	playlists := args.Get(0)
	err := args.Error(1)

	if playlists == nil {
		return nil, err
	}

	return playlists.([]*spotify.Playlist), err
}

func (m *MockAPI) Search(q, searchType string, limit int) (*spotify.Paging, error) {
	args := m.Called(q, searchType, limit)

	paging := args.Get(0)
	err := args.Error(1)

	if paging == nil {
		return nil, err
	}

	return paging.(*spotify.Paging), err
}
