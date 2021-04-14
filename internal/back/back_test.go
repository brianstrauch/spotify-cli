package back

import (
	"errors"
	"spotify/internal"
	"spotify/pkg"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBackCommand(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Back").Return(nil)

	err := back(api)
	require.NoError(t, err)
}

func TestNoPreviousErr(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Back").Return(errors.New(internal.RestrictionViolatedSpotifyErr))

	err := back(api)
	require.Equal(t, internal.NoPreviousErr, err.Error())
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Back").Return(errors.New(internal.NoActiveDeviceSpotifyErr))

	err := back(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
