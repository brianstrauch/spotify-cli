package pkg

import "spotify/pkg/model"

type Token struct {
	AccessToken string `mapstructure:"access_token"`
	ExpiresAt   int64  `mapstructure:"expires_at"`
}

func NewToken(token *model.RelativeToken, now int64) *Token {
	return &Token{
		AccessToken: token.AccessToken,
		ExpiresAt:   now + int64(token.ExpiresIn),
	}
}
