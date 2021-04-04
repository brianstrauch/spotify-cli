package internal

import (
	"os/exec"
	"runtime"
	"spotify/pkg"
	"spotify/pkg/model"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewLoginCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Log in to Spotify.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			token, err := authorize()
			if err != nil {
				return err
			}

			if err := persist(token); err != nil {
				return err
			}

			cmd.Println("Success!")
			return nil
		},
	}
}

func authorize() (*model.Token, error) {
	// https://developer.spotify.com/documentation/general/guides/authorization-guide/#authorization-code-flow-with-proof-key-for-code-exchange-pkce

	// 1. Create the code verifier and challenge
	verifier, challenge := pkg.StartProof()

	// 2. Construct the authorization URI
	uri := pkg.BuildAuthURI(challenge)

	// 3. Your app redirects the user to the authorization URI
	if err := exec.Command(findOpenCommand(), uri).Run(); err != nil {
		return nil, err
	}

	code, err := pkg.ListenForCode()
	if err != nil {
		return nil, err
	}

	// 4. Your app exchanges the code for an access token
	token, err := pkg.RequestToken(code, verifier)
	if err != nil {
		return nil, err
	}

	return token, err
}

func persist(token *model.Token) error {
	// Save token
	viper.Set("token", token.AccessToken)

	// Save expiration
	expiration := time.Now().Unix() + int64(token.ExpiresIn)
	viper.Set("expiration", expiration)

	// Save refresh token
	viper.Set("refresh_token", token.RefreshToken)

	return viper.WriteConfig()
}

func findOpenCommand() string {
	switch os := runtime.GOOS; os {
	case "linux":
		return "xdg-open"
	default:
		return "open"
	}
}
