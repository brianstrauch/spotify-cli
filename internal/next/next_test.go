package next

import (
	"errors"
	"spotify/internal"
	"spotify/pkg"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNextCommand(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Next").Return(nil)

	err := next(api)
	require.NoError(t, err)
}

func TestNoNextErr(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Next").Return(errors.New(internal.RestrictionViolatedSpotifyErr))

	err := next(api)
	require.Equal(t, internal.NoNextErr, err.Error())
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Next").Return(errors.New(internal.NoActiveDeviceSpotifyErr))

	err := next(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
