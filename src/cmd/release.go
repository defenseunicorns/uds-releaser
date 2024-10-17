/*
Copyright © 2024 Defense Unicorns

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/defenseunicorns/uds-releaser/src/platforms"
	"github.com/defenseunicorns/uds-releaser/src/platforms/github"
	"github.com/defenseunicorns/uds-releaser/src/platforms/gitlab"
	"github.com/spf13/cobra"
)

var tokenVarName string

// gitlabCmd represents the gitlab command
var gitlabCmd = &cobra.Command{
	Use:   "gitlab flavor",
	Short: "Create a tag and release on GitLab based on flavor",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return platforms.LoadAndTag(releaserDir, args[0], tokenVarName, gitlab.Platform{})
	},
}

// githubCmd represents the github command
var githubCmd = &cobra.Command{
	Use:   "github flavor",
	Short: "Create a tag and release on GitHub based on flavor",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return platforms.LoadAndTag(releaserDir, args[0], tokenVarName, github.Platform{})
	},
}

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release platform",
	Short: "Collection of commands for releasing on different platforms",
}

func init() {
	rootCmd.AddCommand(releaseCmd)
	releaseCmd.AddCommand(gitlabCmd)
	releaseCmd.AddCommand(githubCmd)
	gitlabCmd.Flags().StringVarP(&tokenVarName, "token-var-name", "t", "GITLAB_RELEASE_TOKEN", "Environment variable name for GitLab token")
	githubCmd.Flags().StringVarP(&tokenVarName, "token-var-name", "t", "GITHUB_TOKEN", "Environment variable name for GitHub token")
}
