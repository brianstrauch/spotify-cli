package internal

import (
	"errors"
	"strings"
	"time"

	"github.com/brianstrauch/spotify"
	"github.com/spf13/viper"
)

const ClientID = "81dddfee3e8d47d89b7902ba616f3357"

func Authenticate() (*spotify.API, error) {
	if time.Now().Unix() > viper.GetInt64("expiration") {
		if err := refresh(); err != nil {
			return nil, err
		}
	}

	token := viper.GetString("token")
	if token == "" {
		return nil, errors.New(ErrNotLoggedIn)
	}

	return spotify.NewAPI(token), nil
}

func refresh() error {
	refresh := viper.GetString("refresh_token")

	token, err := spotify.RefreshPKCEToken(refresh, ClientID)
	if err != nil {
		return err
	}

	return SaveToken(token)
}

func SaveToken(token *spotify.Token) error {
	expiration := time.Now().Unix() + int64(token.ExpiresIn)

	viper.Set("expiration", expiration)
	viper.Set("token", token.AccessToken)
	viper.Set("refresh_token", token.RefreshToken)

	return viper.WriteConfig()
}

func WaitForUpdatedPlayback(api APIInterface, isUpdated func(*spotify.Playback) bool) (*spotify.Playback, error) {
	timeout := time.After(time.Second)
	tick := time.NewTicker(100 * time.Millisecond)

	for {
		select {
		case <-timeout:
			return nil, errors.New(ErrRequestTimedOut)
		case <-tick.C:
			playback, err := api.GetPlayback()
			if err != nil {
				return nil, err
			}

			if isUpdated(playback) {
				return playback, nil
			}
		}
	}
}

func SearchTrack(api APIInterface, query string) (*spotify.Track, error) {
	paging, err := api.Search(query, "track", 1)
	if err != nil {
		return nil, err
	}

	tracks := paging.Tracks.Items
	if len(tracks) == 0 {
		return nil, errors.New(ErrTrackNotFound)
	}

	return paging.Tracks.Items[0], nil
}

func SearchAlbum(api APIInterface, query string) (*spotify.Album, error) {
	paging, err := api.Search(query, "album", 1)
	if err != nil {
		return nil, err
	}

	albums := paging.Albums.Items
	if len(albums) == 0 {
		return nil, errors.New(ErrAlbumNotFound)
	}

	return albums[0], nil
}

func SearchPlaylist(api APIInterface, query string) (*spotify.Playlist, error) {
	playlists, err := api.GetPlaylists()
	if err != nil {
		return nil, err
	}

	for _, playlist := range playlists {
		if strings.EqualFold(playlist.Name, query) {
			return playlist, nil
		}
	}

	return nil, errors.New(ErrPlaylistNotFound)
}
