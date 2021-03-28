package internal

import (
	"os/exec"
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
			token, err := Authorize(cmd)
			if err != nil {
				return err
			}

			// TODO: Group these values under "token"
			viper.Set("token", token.AccessToken)
			viper.Set("expiration", time.Now().Unix()+int64(token.ExpiresIn))

			if err := viper.WriteConfig(); err != nil {
				return err
			}

			cmd.Println("You are logged in!")
			return nil
		},
	}
}

func Authorize(cmd *cobra.Command) (*model.Token, error) {
	// https://developer.spotify.com/documentation/general/guides/authorization-guide/#authorization-code-flow-with-proof-key-for-code-exchange-pkce

	// 1. Create the code verifier and challenge
	verifier, challenge := pkg.StartProof()

	// 2. Construct the authorization URI
	uri := pkg.BuildAuthURI(challenge)

	// 3. Your app redirects the user to the authorization URI
	// TODO: Support other operating systems
	if err := exec.Command("open", uri).Run(); err != nil {
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

func IsAuthenticated() bool {
	exp := viper.GetInt64("expiration")
	now := time.Now().Unix()
	return now < exp
}
