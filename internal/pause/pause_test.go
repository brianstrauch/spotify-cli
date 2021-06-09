package pause

import (
	"errors"
	"spotify/internal"
	"testing"

	"github.com/brianstrauch/spotify"
	"github.com/brianstrauch/spotify/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPauseCommand(t *testing.T) {
	api := new(spotify.MockAPI)

	playback1 := &model.Playback{
		IsPlaying:  true,
		ProgressMs: 0,
		Item: model.Item{
			Type: "track",
			Name: "Song",
			Artists: []model.Artist{
				{Name: "Artist"},
			},
			DurationMs: 1000,
		},
	}

	playback2 := new(model.Playback)
	*playback2 = *playback1
	playback2.IsPlaying = false

	api.On("Status").Return(playback1, nil)
	api.On("WaitForUpdatedPlayback", mock.AnythingOfType("func(*model.Playback) bool")).Return(playback2, nil)
	api.On("Pause").Return(nil)

	status, err := Pause(api)
	require.Equal(t, "🎵 Song\n🎤 Artist\n⏸  0:00 [                ] 0:01\n", status)
	require.NoError(t, err)
}

func TestAlreadyPaused(t *testing.T) {
	api := new(spotify.MockAPI)
	api.On("Status").Return(new(model.Playback), nil)
	api.On("Pause").Return(errors.New(internal.RestrictionViolatedSpotifyErr))

	_, err := Pause(api)
	require.Equal(t, internal.AlreadyPausedErr, err.Error())
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(spotify.MockAPI)
	api.On("Status").Return(nil, nil)

	_, err := Pause(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
