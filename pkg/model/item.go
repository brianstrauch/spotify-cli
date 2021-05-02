package model

type Item struct {
	Artists    []Artist `json:"artists"`
	DurationMs int      `json:"duration_ms"`
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Show       Show     `json:"show"`
	Type       string   `json:"type"`
}
