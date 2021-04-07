package model

type Item struct {
	Name    string   `json:"name"`
	Artists []Artist `json:"artists"`
}
