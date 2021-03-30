package internal

import (
	"os/exec"
	"runtime"
	"spotify/pkg"
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

			// Save token
			viper.Set("token", token)
			if err := viper.WriteConfig(); err != nil {
				return err
			}

			cmd.Println("Success!")
			return nil
		},
	}
}

func authorize() (*pkg.Token, error) {
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
	relativeToken, err := pkg.RequestToken(code, verifier)
	if err != nil {
		return nil, err
	}

	token := pkg.NewToken(relativeToken, time.Now().Unix())

	return token, err
}

func findOpenCommand() string {
	switch os := runtime.GOOS; os {
	case "linux":
		return "xdg-open"
	default:
		return "open"
	}
}
