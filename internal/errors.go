package internal

const (
	ErrAlbumNotFound       = "album not found"
	ErrAlreadyUpToDate     = "already up to date"
	ErrInvalidPlayArgs     = "you may only pass args, --album, or --playlist"
	ErrLoginFailed         = "login failed"
	ErrNoActiveDevice      = "no active spotify session found"
	ErrNoDevices           = "no devices found"
	ErrNoPlaylists         = "no playlists found"
	ErrNoPrevious          = "no track before this one"
	ErrNotLoggedIn         = `you are not logged in, run "spotify login"`
	ErrPlaylistNotFound    = "playlist not found"
	ErrRequestTimedOut     = "request timed out"
	ErrSavePodcast         = "saving podcasts is not allowed"
	ErrTrackNotFound       = "track not found"
)
