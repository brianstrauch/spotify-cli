package save

import (
	"spotify/internal"
	"testing"

	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
)

func TestSaveCommand(t *testing.T) {
	api := new(internal.MockAPI)

	var id string

	playback := &spotify.Playback{
		Item: spotify.Item{
			ID: id,
		},
	}

	api.On("GetPlayback").Return(playback, nil)
	api.On("SaveTracks", []string{id}).Return(nil)

	err := save(api)
	require.NoError(t, err)
}

func TestSavePodcast(t *testing.T) {
	api := new(internal.MockAPI)

	playback := &spotify.Playback{
		Item: spotify.Item{
			Type: "episode",
		},
	}

	api.On("GetPlayback").Return(playback, nil)

	err := save(api)
	require.Equal(t, internal.ErrSavePodcast, err.Error())
}
