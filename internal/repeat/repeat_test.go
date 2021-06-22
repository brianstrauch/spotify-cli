package repeat

import (
	"errors"
	"spotify/internal"
	"testing"

	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
)

func TestRepeatCommandOn(t *testing.T) {
	api := new(internal.MockAPI)

	playback1 := &spotify.Playback{RepeatState: StateOff}
	playback2 := &spotify.Playback{RepeatState: StateOn}

	api.On("GetPlayback").Return(playback1, nil).Once()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("Repeat", StateOn).Return(nil)

	err := Repeat(api, StateOn)
	require.NoError(t, err)
}

func TestRepeatCommandOff(t *testing.T) {
	api := new(internal.MockAPI)

	playback1 := &spotify.Playback{RepeatState: StateOn}
	playback2 := &spotify.Playback{RepeatState: StateOff}

	api.On("GetPlayback").Return(playback1, nil).Once()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("Repeat", StateOff).Return(nil)

	err := Repeat(api, StateOff)
	require.NoError(t, err)
}

func TestRepeatCommandTrack(t *testing.T) {
	api := new(internal.MockAPI)

	playback1 := &spotify.Playback{RepeatState: StateOn}
	playback2 := &spotify.Playback{RepeatState: StateTrack}

	api.On("GetPlayback").Return(playback1, nil).Once()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("Repeat", StateTrack).Return(nil)

	err := Repeat(api, StateTrack)
	require.NoError(t, err)
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(internal.MockAPI)
	api.On("Repeat", StateOn).Return(errors.New(internal.ErrNoActiveDevice))

	err := Repeat(api, StateOn)
	require.Equal(t, internal.ErrNoActiveDevice, err.Error())
}
