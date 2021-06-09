package shuffle

import (
	"spotify/internal"
	"testing"

	"github.com/brianstrauch/spotify"
	"github.com/brianstrauch/spotify/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestShuffleCommandOn(t *testing.T) {
	api := new(spotify.MockAPI)

	playback1 := &model.Playback{ShuffleState: false}
	playback2 := &model.Playback{ShuffleState: true}

	api.On("Status").Return(playback1, nil)
	api.On("WaitForUpdatedPlayback", mock.AnythingOfType("func(*model.Playback) bool")).Return(playback2, nil)
	api.On("Shuffle", true).Return(nil)

	status, err := Shuffle(api)
	require.Equal(t, true, status)
	require.NoError(t, err)
}

func TestShuffleCommandOff(t *testing.T) {
	api := new(spotify.MockAPI)

	playback1 := &model.Playback{ShuffleState: true}
	playback2 := &model.Playback{ShuffleState: false}

	api.On("Status").Return(playback1, nil)
	api.On("WaitForUpdatedPlayback", mock.AnythingOfType("func(*model.Playback) bool")).Return(playback2, nil)
	api.On("Shuffle", false).Return(nil)

	status, err := Shuffle(api)
	require.Equal(t, false, status)
	require.NoError(t, err)
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(spotify.MockAPI)
	api.On("Status").Return(nil, nil)

	_, err := Shuffle(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
