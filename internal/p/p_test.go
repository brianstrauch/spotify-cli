package p

import (
	"spotify/internal"
	"testing"

	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
)

func TestPCommandPlay(t *testing.T) {
	api := new(internal.MockAPI)

	playback1 := &spotify.Playback{
		IsPlaying:  false,
		ProgressMs: 0,
		Item: spotify.Item{
			ID:   "0",
			Type: "track",
			Name: "Song",
			Artists: []spotify.Artist{
				{Name: "Artist"},
			},
			DurationMs: 1000,
		},
	}

	playback2 := new(spotify.Playback)
	*playback2 = *playback1
	playback2.IsPlaying = true

	api.On("GetPlayback").Return(playback1, nil).Twice()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("Play", []string(nil)).Return(nil)

	status, err := p(api, "")
	require.NoError(t, err)
	require.Equal(t, "   Song\r🎵\n   Artist\r🎤\n   0:00 [                ] 0:01\r▶️\n", status)
}

func TestPCommandPause(t *testing.T) {
	api := new(internal.MockAPI)

	playback1 := &spotify.Playback{
		IsPlaying:  true,
		ProgressMs: 0,
		Item: spotify.Item{
			Type: "track",
			Name: "Song",
			Artists: []spotify.Artist{
				{Name: "Artist"},
			},
			DurationMs: 1000,
		},
	}

	playback2 := new(spotify.Playback)
	*playback2 = *playback1
	playback2.IsPlaying = false

	api.On("GetPlayback").Return(playback1, nil).Twice()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("Pause").Return(nil)

	status, err := p(api, "")
	require.NoError(t, err)
	require.Equal(t, "   Song\r🎵\n   Artist\r🎤\n   0:00 [                ] 0:01\r⏸\n", status)
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(internal.MockAPI)
	api.On("GetPlayback").Return(nil, nil)

	_, err := p(api, "")
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
