package play

import (
	"errors"
	"spotify/internal"
	"testing"
	"time"

	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
)

func TestPlay(t *testing.T) {
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

	api.On("GetPlayback").Return(playback1, nil).Once()
	api.On("GetPlayback").Return(playback2, nil).Once()
	api.On("Play", "","", []string(nil)).Return(nil)

	status, err := Play(api, "", "", "")
	require.NoError(t, err)
	require.Equal(t, "   Track\r🎵\n   Artist\r🎤\n   0:00 [                ] 0:01\r▶️\n", status)
}

func TestPlay_WithArgs(t *testing.T) {
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

	api.On("Search", query, "track", 1).Return(paging, nil)
	api.On("Play", "", "", []string{uri}).Return(nil)
	api.On("GetPlayback").Return(playback1, nil).Twice()
	api.On("GetPlayback").Return(playback2, nil).Once()

	status, err := Play(api, query, "", "")
	require.NoError(t, err)
	require.Equal(t, "   Track\r🎵\n   Artist\r🎤\n   0:00 [                ] 0:01\r▶️\n", status)
}

func TestPlay_ErrAlreadyPlaying(t *testing.T) {
	api := new(internal.MockAPI)
	api.On("GetPlayback").Return(new(spotify.Playback), nil)
	api.On("Play", "", "", []string(nil)).Return(errors.New(internal.ErrRestrictionViolated))

	_, err := Play(api, "", "", "")
	require.Error(t, err)
	require.Equal(t, internal.ErrRestrictionViolated, err.Error())
}

func TestPlay_ErrNoActiveDevice(t *testing.T) {
	api := new(internal.MockAPI)
	api.On("GetPlayback").Return(nil, nil)

	_, err := Play(api, "", "", "")
	require.Error(t, err)
	require.Equal(t, internal.ErrNoActiveDevice, err.Error())
}
