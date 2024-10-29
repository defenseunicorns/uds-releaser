package utils

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func DoesTagExist(tag string) (bool, error) {
	repo, err := OpenRepo()
	if err != nil {
		return false, err
	}

	tags, err := repo.Tags()
	if err != nil {
		return false, err
	}

	tagExists := false
	err = tags.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().Short() == tag {
			tagExists = true
		}
		return nil
	})
	return tagExists, err
}

func OpenRepo() (*git.Repository, error) {
	return git.PlainOpen(".")
}

func GetRepoInfo() (remoteURL string, defaultBranch string, err error) {
	repo, err := OpenRepo()
	if err != nil {
		return "", "",  err
	}

	remote, err := repo.Remote("origin")
	if err != nil {
		return "", "", err
	}

	remoteURL = remote.Config().URLs[0]

	ref, err := repo.Head()
	if err != nil {
		return "", "", err
	}

	defaultBranch = ref.Name().Short()
	return remoteURL, defaultBranch, nil
}
