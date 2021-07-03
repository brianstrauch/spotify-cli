package save

import (
	"spotify/internal"
	"testing"

	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {
	api := new(internal.MockAPI)

	var id string

	playback := &spotify.Playback{
		Item: spotify.Item{
			Track: spotify.Track{
				Meta: spotify.Meta{ID: id},
			},
		},
	}

	api.On("GetPlayback").Return(playback, nil)
	api.On("SaveTracks", []string{id}).Return(nil)

	err := save(api)
	require.NoError(t, err)
}

func TestSave_ErrSavePodcast(t *testing.T) {
	api := new(internal.MockAPI)

	playback := &spotify.Playback{
		Item: spotify.Item{
			Type: "episode",
		},
	}

	api.On("GetPlayback").Return(playback, nil)

	err := save(api)
	require.Error(t, err)
	require.Equal(t, internal.ErrSavePodcast, err.Error())
}
