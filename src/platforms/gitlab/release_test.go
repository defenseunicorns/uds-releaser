package gitlab

import (
	"testing"

	"github.com/defenseunicorns/uds-releaser/src/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateReleaseOptions(t *testing.T) {
	packageName := "testing-package"
	flavor := types.Flavor{
		Name:    "unicorn",
		Version: "1.0.0-uds.0",
	}

	defaultBranch := "main"

	releaseOpts := createReleaseOptions(packageName, flavor, defaultBranch)

	assert.Equal(t, "testing-package 1.0.0-uds.0-unicorn", *releaseOpts.Name)
}

func TestGetGitlabBaseUrl(t *testing.T) {
	tests := []struct {
		name     string
		remoteURL string
		expected string
	}{
		{
			name:     "ssh",
			remoteURL: "git@gitlab.com:defenseunicorns/uds-releaser.git",
			expected: "https://gitlab.com/api/v4",
		},
		{
			name:     "https",
			remoteURL: "https://gitlab.com/defenseunicorns/uds-releaser.git",
			expected: "https://gitlab.com/api/v4",
		},
		{
			name:     "ssh-self-hosted",
			remoteURL: "git@gitlab.fake.com:defenseunicorns/uds-releaser.git",
			expected: "https://gitlab.fake.com/api/v4",
		},
		{
			name:     "https-self-hosted",
			remoteURL: "https://gitlab.fake.com/defenseunicorns/uds-releaser.git",
			expected: "https://gitlab.fake.com/api/v4",
		},
		{
			name:     "https-with-token",
			remoteURL: "https://test:token@gitlab.com/defenseunicorns/uds-releaser.git",
			expected: "https://gitlab.com/api/v4",
		},
		{
			name:     "https-self-hosted-with-token",
			remoteURL: "https://test:token@gitlab.fake.com/defenseunicorns/uds-releaser.git",
			expected: "https://gitlab.fake.com/api/v4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, err := getGitlabBaseUrl(tt.remoteURL)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, baseURL)
		})
	}
}
