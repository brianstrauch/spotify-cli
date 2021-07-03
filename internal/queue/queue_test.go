package queue

import (
	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
	"spotify/internal"
	"testing"
)

func TestQueue(t *testing.T) {
	api := new(internal.MockAPI)

	query := "song"
	var uri string

	paging := &spotify.Paging{
		Tracks: spotify.Tracks{
			Items: []spotify.PlaylistTrack{
				{
					URI: uri,
				},
			},
		},
	}

	api.On("Search", query, 1).Return(paging, nil).Once()
	api.On("Queue", uri).Return(nil)

	err := Queue(api, query)
	require.NoError(t, err)
}
