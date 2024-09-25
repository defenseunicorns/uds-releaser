package gitlab

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/defenseunicorns/uds-releaser/src/types"
	"github.com/defenseunicorns/uds-releaser/src/utils"
	gitlab "github.com/xanzy/go-gitlab"
)

func TagAndRelease(flavor types.Flavor) error {
	repo, err := utils.OpenRepo()
	if err != nil {
		return err
	}

	remote, err := repo.Remote("origin")
	if err != nil {
		return err
	}

	remoteURL := remote.Config().URLs[0]

	// Parse the GitLab base URL from the remote URL
	var gitlabBaseURL string
	if strings.Contains(remoteURL, "gitlab.com") {
		gitlabBaseURL = "https://gitlab.com/api/v4"
	} else {
		// Extract the base URL for self-hosted GitLab instances
		parts := strings.Split(remoteURL, "/")
		if len(parts) > 2 {
			gitlabBaseURL = fmt.Sprintf("https://%s/api/v4", parts[2])
		} else {
			return errors.New(fmt.Sprintf("error parsing gitlab base url from remote url: %s", remoteURL))
		}
	}

	// Get the default branch of the current repository
	ref, err := repo.Head()
	if err != nil {
		return err
	}

	defaultBranch := ref.Name().Short()

	// Create a new GitLab client
	gitlabClient, err := gitlab.NewClient(os.Getenv("CI_JOB_TOKEN"), gitlab.WithBaseURL(gitlabBaseURL))
	if err != nil {
		return err
	}

	zarfPackageName, err := utils.GetPackageName()
	if err != nil {
		return err
	}

	// setup the release options
	releaseOpts := &gitlab.CreateReleaseOptions{
		Name:        gitlab.Ptr(fmt.Sprintf("%s %s-%s", zarfPackageName, flavor.Version, flavor.Name)),
		TagName:     gitlab.Ptr(fmt.Sprintf("%s-%s", flavor.Version, flavor.Name)),
		Description: gitlab.Ptr("Release description"),
		Ref:         gitlab.Ptr(defaultBranch),
		Assets: &gitlab.ReleaseAssetsOptions{
			Links: []*gitlab.ReleaseAssetLinkOptions{
				{
					Name: gitlab.Ptr("zarf.yaml"), // TODO
					URL:  gitlab.Ptr("https://example.com/zarf.yaml"), // TODO
					LinkType: gitlab.Ptr(gitlab.PackageLinkType),
				},
				{
					Name: gitlab.Ptr("uds-bundle.yaml"), // TODO
					URL:  gitlab.Ptr("https://example.com/uds-bundle.yaml"), // TODO
				},
			},
		},
	}

	// Create the release
	release, _, err := gitlabClient.Releases.CreateRelease(os.Getenv("CI_PROJECT_ID"), releaseOpts)
	if err != nil {
		return err
	}

	fmt.Printf("Release %s created\n", release.Name)

	return nil
}
