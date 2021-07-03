package internal

const (
	ErrAlreadyPaused       = "Already paused"
	ErrAlreadyPlaying      = "Already playing"
	ErrAlreadyUpToDate     = "Already up to date"
	ErrLoginFailed         = "Login failed"
	ErrNoActiveDevice      = "Player command failed: No active device found"
	ErrNoPlaylists         = "No playlists"
	ErrNoPrevious          = "No track before this one"
	ErrNotLoggedIn         = `You are not logged in: Run "spotify login"`
	ErrRestrictionViolated = "Player command failed: Restriction violated"
	ErrSavePodcast         = "Saving podcasts is not allowed"
)
