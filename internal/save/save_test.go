package save

import (
	"spotify/pkg"
	"spotify/pkg/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveCommand(t *testing.T) {
	api := new(pkg.MockAPI)

	id := ""

	playback := &model.Playback{
		Item: model.Item{
			ID: id,
		},
	}

	api.On("Status").Return(playback, nil)
	api.On("Save", id).Return(nil)

	err := save(api)
	require.NoError(t, err)
}
