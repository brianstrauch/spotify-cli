package version

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersionCommand(t *testing.T) {
	v := version()
	require.True(t, regexp.MustCompile(`^\d+.\d+\.\d+$`).MatchString(v))
}
