package test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckCommand(t *testing.T) {
	stdout, stderr, err := e2e.UDSReleaser("check", "base", "-d", "src/test")
	require.NoError(t, err, stdout, stderr)

	require.Equal(t, "Version 1.0.0-uds.0-base is not tagged\n", stdout)
	
	stdout, stderr, err = e2e.UDSReleaser("check", "dummy", "-d", "src/test")
	require.Error(t, err, stdout, stderr)

	require.Equal(t, "Version testing-dummy is already tagged\n", stdout)
}

func TestCheckCommandBool(t *testing.T) {
	stdout, stderr, err := e2e.UDSReleaser("check", "base", "-d", "src/test", "-b")
	require.NoError(t, err, stdout, stderr)

	require.Equal(t, "true\n", stdout)

	stdout, stderr, err = e2e.UDSReleaser("check", "dummy", "-d", "src/test", "-b")
	require.NoError(t, err, stdout, stderr)

	require.Equal(t, "false\n", stdout)
}
