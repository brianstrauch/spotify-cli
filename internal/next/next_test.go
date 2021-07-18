package next

import (
	"spotify/internal"
	"testing"
	"time"

	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
)

func TestNext(t *testing.T) {
	api := new(internal.MockAPI)

	playback1 := &spotify.Playback{
		IsPlaying:  true,
		ProgressMs: 0,
		Item: spotify.Item{
			Track: spotify.Track{
				Meta:     spotify.Meta{ID: "0"},
				Name:     "Track",
				Artists:  []spotify.Artist{{Name: "Artist"}},
				Duration: &spotify.Duration{Duration: time.Second},
			},
			Type: "track",
		},
	}

	playback2 := new(spotify.Playback)
	*playback2 = *playback1
	playback2.Item.ID = "1"

	api.On("GetPlayback").Return(playback1, nil).Once()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("SkipToNextTrack").Return(nil)

	status, err := next(api)
	require.NoError(t, err)
	require.Equal(t, "   Track\r🎵\n   Artist\r🎤\n   0:00 [                ] 0:01\r▶️\n", status)
}

func TestNext_ErrNoActiveDevice(t *testing.T) {
	api := new(internal.MockAPI)
	api.On("GetPlayback").Return(nil, nil)

	_, err := next(api)
	require.Error(t, err)
	require.Equal(t, internal.ErrNoActiveDevice, err.Error())
}
