package internal

const (
	AlreadyPausedErr              = "Already paused"
	AlreadyPlayingErr             = "Already playing"
	NoActiveDeviceErr             = "Spotify is not active on any device"
	NoActiveDeviceSpotifyErr      = "Player command failed: No active device found"
	NoNextErr                     = "No track after this one"
	NoPreviousErr                 = "No track before this one"
	NotLoggedInErr                = "You are not logged in. Run 'spotify login' before using this command"
	RestrictionViolatedSpotifyErr = "Player command failed: Restriction violated"
	SavePodcastErr                = "Saving podcasts is not allowed"
	TokenExpiredErr               = "API token is expired"
)
