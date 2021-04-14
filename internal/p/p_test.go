package p

import (
	"spotify/internal"
	"spotify/pkg"
	"spotify/pkg/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPCommandPlay(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)

	playback := &model.Playback{
		IsPlaying: false,
	}

	api.On("Status").Return(playback, nil)
	api.On("Play").Return(nil)

	err := p(api)
	require.NoError(t, err)
}

func TestPCommandPause(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)

	playback := &model.Playback{
		IsPlaying: true,
	}

	api.On("Status").Return(playback, nil)
	api.On("Pause").Return(nil)

	err := p(api)
	require.NoError(t, err)
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Status").Return(nil, nil)

	err := p(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
