package gitlab

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/defenseunicorns/uds-releaser/src/types"
	"github.com/defenseunicorns/uds-releaser/src/utils"
	gitlab "github.com/xanzy/go-gitlab"
)

func TagAndRelease(flavor types.Flavor, tokenVarName string) error {
	remoteURL, defaultBranch, _, err := utils.GetRepoInfo()
	if err != nil {
		return err
	}

	// Parse the GitLab base URL from the remote URL
	gitlabBaseURL, err := getGitlabBaseUrl(remoteURL)
	if err != nil {
		return err
	}

	fmt.Printf("Default branch: %s\n", defaultBranch)

	// Create a new GitLab client
	gitlabClient, err := gitlab.NewClient(os.Getenv(tokenVarName), gitlab.WithBaseURL(gitlabBaseURL))
	if err != nil {
		return err
	}

	zarfPackageName, err := utils.GetPackageName()
	if err != nil {
		return err
	}

	// setup the release options
	releaseOpts := createReleaseOptions(zarfPackageName, flavor, defaultBranch)

	fmt.Printf("Creating release %s-%s\n", flavor.Version, flavor.Name)

	// Create the release
	release, _, err := gitlabClient.Releases.CreateRelease(os.Getenv("CI_PROJECT_ID"), releaseOpts)
	if err != nil {
		return err
	}

	fmt.Printf("Release %s created\n", release.Name)

	return nil
}

func createReleaseOptions(zarfPackageName string, flavor types.Flavor, branchRef string) *gitlab.CreateReleaseOptions {
	return &gitlab.CreateReleaseOptions{
		Name:        gitlab.Ptr(fmt.Sprintf("%s %s-%s", zarfPackageName, flavor.Version, flavor.Name)),
		TagName:     gitlab.Ptr(fmt.Sprintf("%s-%s", flavor.Version, flavor.Name)),
		Description: gitlab.Ptr(fmt.Sprintf("%s %s-%s", zarfPackageName, flavor.Version, flavor.Name)),
		Ref:         gitlab.Ptr(branchRef),
	}
}

func getGitlabBaseUrl(remoteURL string) (gitlabBaseURL string, err error) {
	if strings.Contains(remoteURL, "gitlab.com") {
		gitlabBaseURL = "https://gitlab.com/api/v4"
	} else {

		parts := strings.Split(remoteURL, "/")
		containsAt := strings.Contains(remoteURL, "@")
		if len(parts) > 2 {
			gitlabBaseURL = fmt.Sprintf("https://%s/api/v4", parts[2])
		} else if containsAt {
			regex := regexp.MustCompile(`@(.+):`)

			matches := regex.FindStringSubmatch(remoteURL)
			if len(matches) > 1 {
				gitlabBaseURL = fmt.Sprintf("https://%s/api/v4", matches[1])
			} else {
				return "", fmt.Errorf("error parsing gitlab base url from remote url: %s", remoteURL)
			}
		} else {
			return "", fmt.Errorf("error parsing gitlab base url from remote url: %s", remoteURL)
		}
	}

	fmt.Printf("GitLab base URL: %s\n", gitlabBaseURL)
	return gitlabBaseURL, nil
}
