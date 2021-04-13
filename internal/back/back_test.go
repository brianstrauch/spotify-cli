package back

import (
	"spotify/pkg"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBackCommand(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Previous").Return(nil)

	err := back(api)
	require.NoError(t, err)
}
