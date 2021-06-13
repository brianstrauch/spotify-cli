package status

import (
	"spotify/internal"
	"testing"

	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
)

func TestStatusCommand(t *testing.T) {
	api := new(internal.MockAPI)

	playback := &spotify.Playback{
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

	api.On("GetPlayback").Return(playback, nil)

	status, err := status(api)
	require.NoError(t, err)
	require.Equal(t, "   Song\r🎵\n   Artist\r🎤\n   0:00 [                ] 0:01\r▶️\n", status)
}

func TestMultipleArtists(t *testing.T) {
	api := new(internal.MockAPI)

	playback := &spotify.Playback{
		IsPlaying:  true,
		ProgressMs: 0,
		Item: spotify.Item{
			Type: "track",
			Name: "Song",
			Artists: []spotify.Artist{
				{Name: "Artist 1"},
				{Name: "Artist 2"},
			},
			DurationMs: 1000,
		},
	}

	api.On("GetPlayback").Return(playback, nil)

	status, err := status(api)
	require.NoError(t, err)
	require.Equal(t, "   Song\r🎵\n   Artist 1, Artist 2\r🎤\n   0:00 [                ] 0:01\r▶️\n", status)
}

func TestPodcast(t *testing.T) {
	api := new(internal.MockAPI)

	playback := &spotify.Playback{
		IsPlaying:  true,
		ProgressMs: 0,
		Item: spotify.Item{
			Type: "episode",
			Name: "Episode",
			Show: spotify.Show{
				Name: "Podcast",
			},
			DurationMs: 1000,
		},
	}

	api.On("GetPlayback").Return(playback, nil)

	status, err := status(api)
	require.NoError(t, err)
	require.Equal(t, "   Episode\r🎵\n   Podcast\r🎤\n   0:00 [                ] 0:01\r▶️\n", status)
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(internal.MockAPI)
	api.On("GetPlayback").Return(nil, nil)

	_, err := status(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
