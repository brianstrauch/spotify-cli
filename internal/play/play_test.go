package play

import (
	"errors"
	"spotify/internal"
	"spotify/pkg"
	"spotify/pkg/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayCommand(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)

	playback := &model.Playback{
		Item: model.Item{
			Name: "Song",
			Artists: []model.Artist{
				{Name: "Artist"},
			},
		},
	}

	api.On("Status").Return(playback, nil)
	api.On("Play").Return(nil)

	status, err := Play(api)
	require.Equal(t, "Song\nArtist\n", status)
	require.NoError(t, err)
}

func TestAlreadyPlayingErr(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Status").Return(new(model.Playback), nil)
	api.On("Play").Return(errors.New(internal.RestrictionViolatedSpotifyErr))

	_, err := Play(api)
	require.Equal(t, internal.AlreadyPlayingErr, err.Error())
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Status").Return(nil, nil)

	_, err := Play(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
