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

func (m *MockAPI) Play() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockAPI) Save(id string) error {
	args := m.Called(id)
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
