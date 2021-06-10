package repeat

import (
	"spotify/internal"
	"testing"

	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
)

func TestRepeatCommandOn(t *testing.T) {
	api := new(spotify.MockAPI)

	playback1 := &spotify.Playback{RepeatState: StateOff}
	playback2 := &spotify.Playback{RepeatState: StateOn}

	api.On("GetPlayback").Return(playback1, nil).Once()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("Repeat", StateOn).Return(nil)

	status, err := Repeat(api)
	require.NoError(t, err)
	require.Equal(t, StateOn, status)
}

func TestRepeatCommandTrack(t *testing.T) {
	api := new(spotify.MockAPI)

	playback1 := &spotify.Playback{RepeatState: StateOn}
	playback2 := &spotify.Playback{RepeatState: StateTrack}

	api.On("GetPlayback").Return(playback1, nil).Once()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("Repeat", StateTrack).Return(nil)

	status, err := Repeat(api)
	require.NoError(t, err)
	require.Equal(t, StateTrack, status)
}

func TestRepeatCommandOff(t *testing.T) {
	api := new(spotify.MockAPI)

	playback1 := &spotify.Playback{RepeatState: StateTrack}
	playback2 := &spotify.Playback{RepeatState: StateOff}

	api.On("GetPlayback").Return(playback1, nil).Once()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("Repeat", StateOff).Return(nil)

	status, err := Repeat(api)
	require.NoError(t, err)
	require.Equal(t, StateOff, status)
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(spotify.MockAPI)
	api.On("GetPlayback").Return(nil, nil)

	_, err := Repeat(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
