package next

import (
	"errors"
	"spotify/internal"
	"spotify/pkg"
	"spotify/pkg/model"
	"strconv"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNextCommand(t *testing.T) {
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

	i := 0
	api.On("Status").Run(func(_ mock.Arguments) {
		playback.Item.ID = strconv.Itoa(i)
		i++
	}).Return(playback, nil)

	api.On("Next").Return(nil)

	status, err := next(api)
	require.Equal(t, "🎵 Song\n🎤 Artist\n▶️  0:00 [                ] 0:01\n", status)
	require.NoError(t, err)
}

func TestNoNextErr(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Status").Return(new(model.Playback), nil)
	api.On("Next").Return(errors.New(internal.RestrictionViolatedSpotifyErr))

	_, err := next(api)
	require.Equal(t, internal.NoNextErr, err.Error())
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Status").Return(nil, nil)

	_, err := next(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
