package internal

const (
	AlreadyPausedErr              = "Already paused"
	AlreadyPlayingErr             = "Already playing"
	AlreadyUpToDateErr            = "Already up to date"
	NoActiveDeviceErr             = "Spotify is not active on any device"
	NoPreviousErr                 = "No track before this one"
	NotLoggedInErr                = "You are not logged in. Run 'spotify login' before using this command"
	RestrictionViolatedSpotifyErr = "Player command failed: Restriction violated"
	SavePodcastErr                = "Saving podcasts is not allowed"
	LoginFailedErr                = "Login failed"
)
