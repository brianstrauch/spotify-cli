package login

import (
	"context"
	"crypto/rand"
	_ "embed"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"spotify/internal"

	"github.com/brianstrauch/spotify"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var (
	//go:embed success.html
	successHTML string
	//go:embed failure.html
	failureHTML string
)

const RedirectURI = "http://localhost:1024/callback"

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Log in to Spotify.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			token, err := login()
			if err != nil {
				return err
			}

			if err := internal.SaveToken(token); err != nil {
				return err
			}

			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			user, err := api.GetUserProfile()
			if err != nil {
				return err
			}

			cmd.Printf("Logged in as %s.\n", user.DisplayName)
			return nil
		},
	}
}

func login() (*spotify.Token, error) {
	// 1. Create the code verifier and challenge
	verifier, challenge, err := spotify.CreatePKCEVerifierAndChallenge()
	if err != nil {
		return nil, err
	}

	// 2. Construct the authorization URI
	state, err := generateRandomState()
	if err != nil {
		return nil, err
	}

	scopes := []string{
		spotify.ScopeUserLibraryModify,
		spotify.ScopeUserModifyPlaybackState,
		spotify.ScopeUserReadPlaybackState,
	}

	uri := spotify.BuildPKCEAuthURI(internal.ClientID, RedirectURI, challenge, state, scopes...)

	// 3. Your app redirects the user to the authorization URI
	if err := browser.OpenURL(uri); err != nil {
		return nil, err
	}

	code, err := listenForCode(state)
	if err != nil {
		return nil, err
	}

	// 4. Your app exchanges the code for an access token
	token, err := spotify.RequestPKCEToken(internal.ClientID, code, RedirectURI, verifier)
	if err != nil {
		return nil, err
	}

	return token, err
}

func generateRandomState() (string, error) {
	buf := make([]byte, 7)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}

	state := hex.EncodeToString(buf)
	return state, nil
}

func listenForCode(state string) (code string, err error) {
	server := &http.Server{Addr: ":1024"}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state || r.URL.Query().Get("error") != "" {
			err = errors.New(internal.ErrLoginFailed)
			fmt.Fprintln(w, failureHTML)
		} else {
			code = r.URL.Query().Get("code")
			fmt.Fprintln(w, successHTML)
		}

		// Use a separate thread so browser doesn't show a "No Connection" message
		go func() {
			 _ = server.Shutdown(context.TODO())
		}()
	})

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		return "", err
	}

	return
}
