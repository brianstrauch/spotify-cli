package play

import (
	"spotify/pkg"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayCommand(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Play").Return(nil)

	err := play(api)
	require.NoError(t, err)
}
