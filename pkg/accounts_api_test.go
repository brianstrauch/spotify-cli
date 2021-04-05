package pkg

import (
	"net/url"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStartProof(t *testing.T) {
	verifier, challenge, err := StartProof()
	require.NoError(t, err)

	require.True(t, regexp.MustCompile(`^[[:alnum:]_.\-~]{128}$`).MatchString(verifier))

	// Hash with SHA-256 (64 chars)
	// Convert to Base64 (44 chars)
	// Remove trailing = (43 chars)
	require.True(t, regexp.MustCompile(`^[[:alnum:]\-_]{43}$`).MatchString(challenge))
}

func TestBuildAuthURI(t *testing.T) {
	redirectURI := "http://example.com"
	challenge := "challenge"
	state := "state"

	uri := BuildAuthURI(redirectURI, challenge, state)

	require.True(t, strings.Contains(uri, "client_id="+ClientID))
	require.True(t, strings.Contains(uri, "response_type=code"))
	require.True(t, strings.Contains(uri, "redirect_uri="+url.QueryEscape(redirectURI)))
	require.True(t, strings.Contains(uri, "code_challenge_method=S256"))
	require.True(t, strings.Contains(uri, "code_challenge="+challenge))
	require.True(t, strings.Contains(uri, "state="+state))
	require.True(t, strings.Contains(uri, "scope=user-modify-playback-state"))
}
