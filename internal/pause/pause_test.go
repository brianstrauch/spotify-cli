package pause

import (
	"errors"
	"spotify/internal"
	"testing"
	"time"

	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
)

func TestPause(t *testing.T) {
	api := new(internal.MockAPI)

	playback1 := &spotify.Playback{
		IsPlaying:  true,
		ProgressMs: 0,
		Item: spotify.Item{
			Track: spotify.Track{
				Name: "Song",
				Artists: []spotify.Artist{
					{Name: "Artist"},
				},
				Duration: &spotify.Duration{Duration: time.Second},
			},
			Type: "track",
		},
	}

	playback2 := new(spotify.Playback)
	*playback2 = *playback1
	playback2.IsPlaying = false

	api.On("GetPlayback").Return(playback1, nil).Once()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("Pause").Return(nil)

	status, err := Pause(api)
	require.NoError(t, err)
	require.Equal(t, "   Song\r🎵\n   Artist\r🎤\n   0:00 [                ] 0:01\r⏸\n", status)
}

func TestPause_ErrAlreadyPaused(t *testing.T) {
	api := new(internal.MockAPI)
	api.On("GetPlayback").Return(new(spotify.Playback), nil)
	api.On("Pause").Return(errors.New(internal.ErrRestrictionViolated))

	_, err := Pause(api)
	require.Error(t, err)
	require.Equal(t, internal.ErrAlreadyPaused, err.Error())
}

func TestPause_ErrNoActiveDevice(t *testing.T) {
	api := new(internal.MockAPI)
	api.On("GetPlayback").Return(nil, nil)

	_, err := Pause(api)
	require.Error(t, err)
	require.Equal(t, internal.ErrNoActiveDevice, err.Error())
}
