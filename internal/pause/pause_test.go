package pause

import (
	"errors"
	"spotify/internal"
	"spotify/pkg"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPauseCommand(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Pause").Return(nil)

	err := pause(api)
	require.NoError(t, err)
}

func TestAlreadyPaused(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Pause").Return(errors.New(internal.RestrictionViolatedSpotifyErr))

	err := pause(api)
	require.Equal(t, internal.AlreadyPausedErr, err.Error())
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Pause").Return(errors.New(internal.NoActiveDeviceSpotifyErr))

	err := pause(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
