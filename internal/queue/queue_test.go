package queue

import (
	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
	"spotify/internal"
	"testing"
)

func TestQueue(t *testing.T) {
	api := new(internal.MockAPI)

	uri := "uri"

	paging := &spotify.Paging{
		Tracks: spotify.TrackPage{
			Items: []*spotify.Track{
				{
					Meta:    spotify.Meta{URI: uri},
					Name:    "Track",
					Artists: []spotify.Artist{{Name: "Artist"}},
				},
			},
		},
	}

	query := "track"

	api.On("Search", query, "track", 1).Return(paging, nil).Once()
	api.On("Queue", uri).Return(nil)

	output, err := Queue(api, query)
	require.NoError(t, err)
	require.Equal(t, "   Track\r🎵\n   Artist\r🎤\n", output)
}
