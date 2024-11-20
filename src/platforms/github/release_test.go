// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGitHubTag(t *testing.T) {
	tests := []struct {
		name            string
		tagName         string
		releaseName     string
		hash            string
	}{
		{
			name:            "ValidTag",
			tagName:         "v1.0.0-uds.0-unicorn",
			releaseName:     "testing-package v1.0.0-uds.0-unicorn",
			hash:            "1234567890",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag := createGitHubTag(tt.tagName, tt.releaseName, tt.hash)
			assert.Equal(t, tt.tagName, *tag.Tag)
			assert.Equal(t, tt.releaseName, *tag.Message)
			assert.Equal(t, tt.hash, *tag.Object.SHA)
		})
	}
}

func TestGetGithubOwnerAndRepo(t *testing.T) {
	tests := []struct {
		name          string
		remoteURL     string
		expectedOwner string
		expectedRepo  string
		expectError   bool
	}{
		{
			name:          "HTTPSRemoteURL",
			remoteURL:     "https://github.com/defenseunicorns/uds-releaser.git",
			expectedOwner: "defenseunicorns",
			expectedRepo:  "uds-releaser",
			expectError:   false,
		},
		{
			name:          "HTTPSRemoteURLNoGit",
			remoteURL:     "https://github.com/defenseunicorns/uds-releaser",
			expectedOwner: "defenseunicorns",
			expectedRepo:  "uds-releaser",
			expectError:   false,
		},
		{
			name:          "HTTPSRemoteURLWithToken",
			remoteURL:     "https://test:token@github.com/defenseunicorns/uds-releaser.git",
			expectedOwner: "defenseunicorns",
			expectedRepo:  "uds-releaser",
			expectError:   false,
		},
		{
			name:          "SSHRemoteURL",
			remoteURL:     "git@github.com:defenseunicorns/uds-releaser.git",
			expectedOwner: "defenseunicorns",
			expectedRepo:  "uds-releaser",
			expectError:   false,
		},
		{
			name:          "GitlabRemoteURL",
			remoteURL:     "https://gitlab.com/defenseunicorns/uds-releaser.git",
			expectedOwner: "",
			expectedRepo:  "",
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			owner, repo, err := getGithubOwnerAndRepo(tt.remoteURL)
			if tt.expectError {
				assert.Error(t, err)
				assert.Empty(t, owner)
				assert.Empty(t, repo)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOwner, owner)
				assert.Equal(t, tt.expectedRepo, repo)
			}
		})
	}
}
