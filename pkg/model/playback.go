package model

type Playback struct {
	IsPlaying bool `json:"is_playing"`
	Item      Item `json:"item"`
}
