package p

import (
	"spotify/internal"
	"spotify/pkg"
	"spotify/pkg/model"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPCommandPlay(t *testing.T) {
	api := new(pkg.MockAPI)

	playback := &model.Playback{
		IsPlaying:  false,
		ProgressMs: 0,
		Item: model.Item{
			Name: "Song",
			Artists: []model.Artist{
				{Name: "Artist"},
			},
			DurationMs: 1000,
		},
	}

	i := 0
	api.On("Status").Run(func(_ mock.Arguments) {
		if i == 1 {
			playback.IsPlaying = true
		}
		i++
	}).Return(playback, nil)

	api.On("Play").Return(nil)

	status, err := p(api)
	require.Equal(t, "🎵 Song\n🎤 Artist\n▶️  0:00 [                ] 0:01\n", status)
	require.NoError(t, err)
}

func TestPCommandPause(t *testing.T) {
	api := new(pkg.MockAPI)

	playback := &model.Playback{
		IsPlaying:  true,
		ProgressMs: 0,
		Item: model.Item{
			Name: "Song",
			Artists: []model.Artist{
				{Name: "Artist"},
			},
			DurationMs: 1000,
		},
	}

	i := 0
	api.On("Status").Run(func(_ mock.Arguments) {
		if i == 1 {
			playback.IsPlaying = false
		}
		i++
	}).Return(playback, nil)

	api.On("Pause").Return(nil)

	status, err := p(api)
	require.Equal(t, "🎵 Song\n🎤 Artist\n⏸  0:00 [                ] 0:01\n", status)
	require.NoError(t, err)
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(pkg.MockAPI)
	api.On("Status").Return(nil, nil)

	_, err := p(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
