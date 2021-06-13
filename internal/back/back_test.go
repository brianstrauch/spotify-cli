package back

import (
	"errors"
	"spotify/internal"
	"testing"

	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
)

func TestBackCommand(t *testing.T) {
	api := new(internal.MockAPI)

	playback1 := &spotify.Playback{
		IsPlaying:  true,
		ProgressMs: 0,
		Item: spotify.Item{
			ID:   "1",
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
	playback2.Item.ID = "0"

	api.On("GetPlayback").Return(playback1, nil).Once()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("SkipToPreviousTrack").Return(nil)

	status, err := back(api)
	require.NoError(t, err)
	require.Equal(t, "   Song\r🎵\n   Artist\r🎤\n   0:00 [                ] 0:01\r▶️\n", status)
}

func TestNoPreviousErr(t *testing.T) {
	api := new(internal.MockAPI)
	api.On("GetPlayback").Return(new(spotify.Playback), nil)
	api.On("SkipToPreviousTrack").Return(errors.New(internal.RestrictionViolatedSpotifyErr))

	_, err := back(api)
	require.Equal(t, internal.NoPreviousErr, err.Error())
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(internal.MockAPI)
	api.On("GetPlayback").Return(nil, nil)

	_, err := back(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
