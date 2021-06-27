package shuffle

import (
	"errors"
	"spotify/internal"
	"testing"

	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
)

func TestShuffle_On(t *testing.T) {
	api := new(internal.MockAPI)

	playback1 := &spotify.Playback{ShuffleState: false}
	playback2 := &spotify.Playback{ShuffleState: true}

	api.On("GetPlayback").Return(playback1, nil).Once()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("Shuffle", true).Return(nil)

	err := Shuffle(api, true)
	require.NoError(t, err)
}

func TestShuffle_Off(t *testing.T) {
	api := new(internal.MockAPI)

	playback1 := &spotify.Playback{ShuffleState: true}
	playback2 := &spotify.Playback{ShuffleState: false}

	api.On("GetPlayback").Return(playback1, nil).Once()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("Shuffle", false).Return(nil)

	err := Shuffle(api, false)
	require.NoError(t, err)
}

func TestShuffle_ErrNoActiveDevice(t *testing.T) {
	api := new(internal.MockAPI)
	api.On("Shuffle", true).Return(errors.New(internal.ErrNoActiveDevice))

	err := Shuffle(api, true)
	require.Error(t, err)
	require.Equal(t, internal.ErrNoActiveDevice, err.Error())
}
