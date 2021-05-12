package shuffle

import (
	"spotify/internal"
	"spotify/pkg"
	"spotify/pkg/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShuffleCommand(t *testing.T) {
	api := new(pkg.MockAPI)

	playback1 := &model.Playback{ShuffleState: false}
	playback2 := &model.Playback{ShuffleState: true}

	api.On("Status").Return(playback1, nil).Once()
	api.On("Status").Return(playback2, nil).Once()
	api.On("Shuffle", true).Return(nil)

	status, err := Shuffle(api)
	require.Equal(t, true, status)
	require.NoError(t, err)
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(pkg.MockAPI)
	api.On("Status").Return(nil, nil)

	_, err := Shuffle(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
