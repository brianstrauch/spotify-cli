package play

import (
	"errors"
	"spotify/internal"
	"spotify/pkg"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayCommand(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Play").Return(nil)

	err := Play(api)
	require.NoError(t, err)
}

func TestAlreadyPlayingErr(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Play").Return(errors.New(internal.RestrictionViolatedSpotifyErr))

	err := Play(api)
	require.Equal(t, internal.AlreadyPlayingErr, err.Error())
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Play").Return(errors.New(internal.NoActiveDeviceSpotifyErr))

	err := Play(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
