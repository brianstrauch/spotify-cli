package internal

import (
	"bytes"
	"io/ioutil"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersionCommand(t *testing.T) {
	cmd := NewVersionCommand()

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)

	err := cmd.Execute()
	require.NoError(t, err)

	out, err := ioutil.ReadAll(buf)
	require.NoError(t, err)
	require.True(t, regexp.MustCompile(`^\d+.\d+\.\d+\n$`).Match(out))
}
