package repeat

import (
	"spotify/internal"
	"testing"

	"github.com/brianstrauch/spotify"
	"github.com/brianstrauch/spotify/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRepeatCommandOn(t *testing.T) {
	api := new(spotify.MockAPI)

	playback1 := &model.Playback{RepeatState: StateOff}
	playback2 := &model.Playback{RepeatState: StateOn}

	api.On("Status").Return(playback1, nil)
	api.On("WaitForUpdatedPlayback", mock.AnythingOfType("func(*model.Playback) bool")).Return(playback2, nil)
	api.On("Repeat", StateOn).Return(nil)

	status, err := Repeat(api)
	require.Equal(t, StateOn, status)
	require.NoError(t, err)
}

func TestRepeatCommandTrack(t *testing.T) {
	api := new(spotify.MockAPI)

	playback1 := &model.Playback{RepeatState: StateOn}
	playback2 := &model.Playback{RepeatState: StateTrack}

	api.On("Status").Return(playback1, nil)
	api.On("WaitForUpdatedPlayback", mock.AnythingOfType("func(*model.Playback) bool")).Return(playback2, nil)
	api.On("Repeat", StateTrack).Return(nil)

	status, err := Repeat(api)
	require.Equal(t, StateTrack, status)
	require.NoError(t, err)
}

func TestRepeatCommandOff(t *testing.T) {
	api := new(spotify.MockAPI)

	playback1 := &model.Playback{RepeatState: StateTrack}
	playback2 := &model.Playback{RepeatState: StateOff}

	api.On("Status").Return(playback1, nil)
	api.On("WaitForUpdatedPlayback", mock.AnythingOfType("func(*model.Playback) bool")).Return(playback2, nil)
	api.On("Repeat", StateOff).Return(nil)

	status, err := Repeat(api)
	require.Equal(t, StateOff, status)
	require.NoError(t, err)
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(spotify.MockAPI)
	api.On("Status").Return(nil, nil)

	_, err := Repeat(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
