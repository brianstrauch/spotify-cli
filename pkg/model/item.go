package model

type Item struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Artists    []Artist `json:"artists"`
	DurationMs int      `json:"duration_ms"`
}
