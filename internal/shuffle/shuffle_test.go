package shuffle

import (
	"spotify/internal"
	"testing"

	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
)

func TestShuffleCommandOn(t *testing.T) {
	api := new(internal.MockAPI)

	playback1 := &spotify.Playback{ShuffleState: false}
	playback2 := &spotify.Playback{ShuffleState: true}

	api.On("GetPlayback").Return(playback1, nil).Once()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("Shuffle", true).Return(nil)

	status, err := Shuffle(api)
	require.NoError(t, err)
	require.Equal(t, true, status)
}

func TestShuffleCommandOff(t *testing.T) {
	api := new(internal.MockAPI)

	playback1 := &spotify.Playback{ShuffleState: true}
	playback2 := &spotify.Playback{ShuffleState: false}

	api.On("GetPlayback").Return(playback1, nil).Once()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("Shuffle", false).Return(nil)

	status, err := Shuffle(api)
	require.NoError(t, err)
	require.Equal(t, false, status)
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(internal.MockAPI)
	api.On("GetPlayback").Return(nil, nil)

	_, err := Shuffle(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
