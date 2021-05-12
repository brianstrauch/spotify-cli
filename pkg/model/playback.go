package model

type Playback struct {
	IsPlaying    bool   `json:"is_playing"`
	Item         Item   `json:"item"`
	ProgressMs   int    `json:"progress_ms"`
	RepeatState  string `json:"repeat_state"`
	ShuffleState bool   `json:"shuffle_state"`
}
