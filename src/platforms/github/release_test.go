package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGitHubTag(t *testing.T) {
	// Create a new tag
	tagName := "v1.0.0-uds.0-unicorn"
	releaseName := "testing-package v1.0.0-uds.0-unicorn"
	hash := "1234567890"

	tag := createGitHubTag(tagName, releaseName, hash)

	assert.Equal(t, tagName, *tag.Tag)
	assert.Equal(t, releaseName, *tag.Message)
	assert.Equal(t, hash, *tag.Object.SHA)
}

func TestGetGithubOwnerAndRepo(t *testing.T) {
	// Get the owner and repo from a remote URL
	httpsRemoteURL := "https://github.com/defenseunicorns/uds-releaser.git"
	sshRemoteURL := "git@github.com:defenseunicorns/uds-releaser.git"

	owner, repo, err := getGithubOwnerAndRepo(httpsRemoteURL)
	assert.NoError(t, err)
	assert.Equal(t, "defenseunicorns", owner)
	assert.Equal(t, "uds-releaser", repo)

	owner, repo, err = getGithubOwnerAndRepo(sshRemoteURL)
	assert.NoError(t, err)
	assert.Equal(t, "defenseunicorns", owner)
	assert.Equal(t, "uds-releaser", repo)

	gitlabRemoteURL := "https://gitlab.com/defenseunicorns/uds-releaser.git"
	owner, repo, err = getGithubOwnerAndRepo(gitlabRemoteURL)
	assert.Error(t, err)
	assert.Empty(t, owner)
	assert.Empty(t, repo)
}
