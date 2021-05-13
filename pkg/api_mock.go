package pkg

import (
	"spotify/pkg/model"

	"github.com/stretchr/testify/mock"
)

type MockAPI struct {
	mock.Mock
}

func (m *MockAPI) Back() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockAPI) Next() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockAPI) Pause() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockAPI) Play(uri string) error {
	args := m.Called(uri)
	return args.Error(0)
}

func (m *MockAPI) Repeat(state string) error {
	args := m.Called(state)
	return args.Error(0)
}

func (m *MockAPI) Save(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockAPI) Search(queue string, limit int) (*model.Page, error) {
	args := m.Called(queue, limit)

	page := args.Get(0)
	err := args.Error(1)

	if page == nil {
		return nil, err
	}

	return page.(*model.Page), err
}

func (m *MockAPI) Shuffle(state bool) error {
	args := m.Called(state)
	return args.Error(0)
}

func (m *MockAPI) Status() (*model.Playback, error) {
	args := m.Called()

	playback := args.Get(0)
	err := args.Error(1)

	if playback == nil {
		return nil, err
	}

	return playback.(*model.Playback), err
}

func (m *MockAPI) Unsave(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockAPI) WaitForUpdatedPlayback(isUpdated func(playback *model.Playback) bool) (*model.Playback, error) {
	args := m.Called(isUpdated)

	playback := args.Get(0)
	err := args.Error(1)

	if playback == nil {
		return nil, err
	}

	return playback.(*model.Playback), err
}
