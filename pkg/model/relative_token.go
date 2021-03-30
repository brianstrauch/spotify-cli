package model

type RelativeToken struct {
	AccessToken string `json:"access_token"`
	// TODO: token_type
	ExpiresIn int `json:"expires_in"`
	// TODO: refresh_token
}
