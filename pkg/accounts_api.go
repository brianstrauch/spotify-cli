package pkg

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"spotify/pkg/model"
	"strings"
)

const (
	ClientID = "81dddfee3e8d47d89b7902ba616f3357"
	BaseURL  = "https://accounts.spotify.com"
	Port     = 1024
)

func StartProof() (string, string) {
	verifier := generateRandomVerifier()
	hash := sha256.Sum256(verifier)
	challenge := base64.URLEncoding.EncodeToString(hash[:])
	challenge = strings.TrimRight(challenge, "=")

	return string(verifier), challenge
}

func generateRandomVerifier() []byte {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_.-~"

	verifier := make([]byte, 128)
	for i := 0; i < len(verifier); i++ {
		// TODO: Use crypto/rand
		idx := rand.Intn(len(chars))
		verifier[i] = chars[idx]
	}
	return verifier
}

func BuildAuthURI(challenge string) string {
	q := url.Values{}
	q.Add("client_id", ClientID)
	q.Add("response_type", "code")
	q.Add("redirect_uri", buildRedirectURI())
	q.Add("code_challenge_method", "S256")
	q.Add("code_challenge", challenge)
	// TODO: state
	q.Add("scope", "user-modify-playback-state")

	return BaseURL + "/authorize?" + q.Encode()
}

func ListenForCode() (string, error) {
	server := &http.Server{Addr: fmt.Sprintf(":%d", Port)}
	var code string

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code = r.URL.Query().Get("code")

		// TODO: Check for "access denied" (user cancels)
		// TODO: Return full webpage with title
		fmt.Fprintln(w, "Success!")

		// Use a separate thread so browser does not show a "No Connection" message
		go func() {
			server.Shutdown(context.TODO())
		}()
	})

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		return "", err
	}

	return code, nil
}

func RequestToken(code, verifier string) (*model.Token, error) {
	v := url.Values{}
	v.Set("client_id", ClientID)
	v.Set("grant_type", "authorization_code")
	v.Set("code", code)
	v.Set("redirect_uri", buildRedirectURI())
	v.Set("code_verifier", verifier)
	body := strings.NewReader(v.Encode())

	res, err := http.Post(BaseURL+"/api/token", "application/x-www-form-urlencoded", body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// TODO: Handle errors

	token := new(model.Token)
	if err := json.NewDecoder(res.Body).Decode(token); err != nil {
		return nil, err
	}

	return token, nil
}

func RefreshToken(refreshToken string) (*model.Token, error) {
	v := url.Values{}
	v.Set("grant_type", "refresh_token")
	v.Set("refresh_token", refreshToken)
	v.Set("client_id", ClientID)
	body := strings.NewReader(v.Encode())

	res, err := http.Post(BaseURL+"/api/token", "application/x-www-form-urlencoded", body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// TODO: Handle errors

	token := new(model.Token)
	if err := json.NewDecoder(res.Body).Decode(token); err != nil {
		return nil, err
	}

	return token, nil
}

func buildRedirectURI() string {
	return fmt.Sprintf("http://localhost:%d/callback", Port)
}
