package unsave

import (
	"spotify/internal"
	"spotify/pkg"
	"spotify/pkg/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnsaveCommand(t *testing.T) {
	api := new(pkg.MockAPI)

	id := ""

	playback := &model.Playback{
		Item: model.Item{
			ID: id,
		},
	}

	api.On("Status").Return(playback, nil)
	api.On("Unsave", id).Return(nil)

	err := unsave(api)
	require.NoError(t, err)
}

func TestSavePodcast(t *testing.T) {
	api := new(pkg.MockAPI)

	playback := &model.Playback{
		Item: model.Item{
			Type: "episode",
		},
	}

	api.On("Status").Return(playback, nil)

	err := unsave(api)
	require.Equal(t, internal.SavePodcastErr, err.Error())
}
