package repeat

import (
	"spotify/internal"
	"spotify/pkg"
	"spotify/pkg/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepeatCommandOn(t *testing.T) {
	api := new(pkg.MockAPI)

	playback1 := &model.Playback{RepeatState: "off"}
	playback2 := &model.Playback{RepeatState: "context"}

	api.On("Status").Return(playback1, nil).Once()
	api.On("Status").Return(playback2, nil)
	api.On("Repeat", "context").Return(nil)

	status, err := Repeat(api)
	require.Equal(t, "context", status)
	require.NoError(t, err)
}

func TestRepeatCommandTrack(t *testing.T) {
	api := new(pkg.MockAPI)

	playback1 := &model.Playback{RepeatState: "context"}
	playback2 := &model.Playback{RepeatState: "track"}

	api.On("Status").Return(playback1, nil).Once()
	api.On("Status").Return(playback2, nil)
	api.On("Repeat", "track").Return(nil)

	status, err := Repeat(api)
	require.Equal(t, "track", status)
	require.NoError(t, err)
}

func TestRepeatCommandOff(t *testing.T) {
	api := new(pkg.MockAPI)

	playback1 := &model.Playback{RepeatState: "track"}
	playback2 := &model.Playback{RepeatState: "off"}

	api.On("Status").Return(playback1, nil).Once()
	api.On("Status").Return(playback2, nil)
	api.On("Repeat", "off").Return(nil)

	status, err := Repeat(api)
	require.Equal(t, "off", status)
	require.NoError(t, err)
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(pkg.MockAPI)
	api.On("Status").Return(nil, nil)

	_, err := Repeat(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
