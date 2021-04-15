package status

import (
	"spotify/internal"
	"spotify/pkg"
	"spotify/pkg/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatusCommand(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)

	playback := &model.Playback{
		IsPlaying:  true,
		ProgressMs: 0,
		Item: model.Item{
			Name: "Song",
			Artists: []model.Artist{
				{Name: "Artist"},
			},
			DurationMs: 1000,
		},
	}

	api.On("Status").Return(playback, nil)

	status, err := status(api)
	require.Equal(t, "🎵 Song\n🎤 Artist\n▶️  0:00 [                ] 0:01\n", status)
	require.NoError(t, err)
}

func TestMultipleArtists(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)

	playback := &model.Playback{
		IsPlaying:  true,
		ProgressMs: 0,
		Item: model.Item{
			Name: "Song",
			Artists: []model.Artist{
				{Name: "Artist 1"},
				{Name: "Artist 2"},
			},
			DurationMs: 1000,
		},
	}

	api.On("Status").Return(playback, nil)

	status, err := status(api)
	require.Equal(t, "🎵 Song\n🎤 Artist 1, Artist 2\n▶️  0:00 [                ] 0:01\n", status)
	require.NoError(t, err)
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Status").Return(nil, nil)

	_, err := status(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
