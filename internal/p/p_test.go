package p

import (
	"spotify/internal"
	"testing"
	"time"

	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
)

func TestP_Play(t *testing.T) {
	api := new(internal.MockAPI)

	playback1 := &spotify.Playback{
		IsPlaying:  false,
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
	playback2.IsPlaying = true

	api.On("GetPlayback").Return(playback1, nil).Twice()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("Play", "", []string(nil)).Return(nil)

	status, err := p(api, "", "")
	require.NoError(t, err)
	require.Equal(t, "   Track\r🎵\n   Artist\r🎤\n   0:00 [                ] 0:01\r▶️\n", status)
}

func TestP_Play_WithArgs(t *testing.T) {
	api := new(internal.MockAPI)

	uri := "uri"
	name := "Track"

	paging := &spotify.Paging{
		Tracks: spotify.TrackPage{
			Items: []*spotify.Track{
				{
					Meta:    spotify.Meta{URI: uri},
					Name:    name,
					Artists: []spotify.Artist{{Name: "Artist"}},
				},
			},
		},
	}

	playback1 := &spotify.Playback{}
	playback2 := &spotify.Playback{
		IsPlaying:  true,
		ProgressMs: 0,
		Item: spotify.Item{
			Track: spotify.Track{
				Meta:     spotify.Meta{ID: "0"},
				Name:     name,
				Artists:  []spotify.Artist{{Name: "Artist"}},
				Duration: &spotify.Duration{Duration: time.Second},
			},
			Type: "track",
		},
	}

	query := "track"

	api.On("Search", query, 1).Return(paging, nil)
	api.On("Play", "", []string{uri}).Return(nil)
	api.On("GetPlayback").Return(playback1, nil).Twice()
	api.On("GetPlayback").Return(playback2, nil).Once()

	status, err := p(api, query, "")
	require.NoError(t, err)
	require.Equal(t, "   Track\r🎵\n   Artist\r🎤\n   0:00 [                ] 0:01\r▶️\n", status)
}

func TestP_Pause(t *testing.T) {
	api := new(internal.MockAPI)

	playback1 := &spotify.Playback{
		IsPlaying:  true,
		ProgressMs: 0,
		Item: spotify.Item{
			Track: spotify.Track{
				Name:     "Track",
				Artists:  []spotify.Artist{{Name: "Artist"}},
				Duration: &spotify.Duration{Duration: time.Second},
			},
			Type: "track",
		},
	}

	playback2 := new(spotify.Playback)
	*playback2 = *playback1
	playback2.IsPlaying = false

	api.On("GetPlayback").Return(playback1, nil).Twice()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("Pause", "").Return(nil)

	status, err := p(api, "", "")
	require.NoError(t, err)
	require.Equal(t, "   Track\r🎵\n   Artist\r🎤\n   0:00 [                ] 0:01\r⏸\n", status)
}

func TestP_ErrNoActiveDevice(t *testing.T) {
	api := new(internal.MockAPI)
	api.On("GetPlayback").Return(nil, nil)

	_, err := p(api, "", "")
	require.Error(t, err)
	require.Equal(t, internal.ErrNoActiveDevice, err.Error())
}
