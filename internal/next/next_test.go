package next

import (
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
