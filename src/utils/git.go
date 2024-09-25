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
	if err != nil {
		return tagExists, err
	}
	return tagExists, nil
}

func OpenRepo() (*git.Repository, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return nil, err
	}
	return repo, nil
}