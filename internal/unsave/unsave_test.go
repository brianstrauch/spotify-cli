package unsave

import (
	"spotify/internal"
	"testing"

	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
)

func TestUnsave(t *testing.T) {
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
	api.On("RemoveSavedTracks", []string{id}).Return(nil)

	err := unsave(api)
	require.NoError(t, err)
}

func TestUnsave_ErrSavePodcast(t *testing.T) {
	api := new(internal.MockAPI)

	playback := &spotify.Playback{Item: spotify.Item{Track: spotify.Track{Meta: spotify.Meta{Type: "episode"}}}}

	api.On("GetPlayback").Return(playback, nil)

	err := unsave(api)
	require.Error(t, err)
	require.Equal(t, internal.ErrSavePodcast, err.Error())
}
