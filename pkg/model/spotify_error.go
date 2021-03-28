package model

type SpotifyError struct {
	Error struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Reason  string `json:"reason"`
	} `json:"error"`
}
