package commands

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	argocdclient "github.com/argoproj/argo-cd/v3/pkg/apiclient"
)

func TestNewUpgradeCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := NewUpgradeCmd(&argocdclient.ClientOptions{})
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"upgrade"})
	require.NoError(t, cmd.Execute(), "Failed to execute short version command")
	assert.Equal(t, "argocd: v99.99.99+unknown\n", buf.String())
}
