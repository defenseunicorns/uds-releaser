/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/defenseunicorns/uds-releaser/src/utils"
	"github.com/defenseunicorns/uds-releaser/src/version"
	"github.com/spf13/cobra"
)

// gitlabCmd represents the gitlab command
var gitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "Collection of commands for releasing on GitLab",
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version [ flavor ]",
	Short: "Mutate version fields in the zarf.yaml and uds-bundle.yaml based on flavor",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		releaserConfig, err := utils.LoadReleaserConfig()
		if err != nil {
			return err
		}

		currentFlavor, err := utils.GetFlavorConfig(args[0], releaserConfig)
		if err != nil {
			return err
		}

		err = version.MutateYamls(currentFlavor)
		if err != nil {
			return err
		}
		return nil
	},
}

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release [ flavor ]",
	Short: "Create a tag and release on GitLab based on flavor",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("release called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(gitlabCmd)
	gitlabCmd.AddCommand(versionCmd)
	gitlabCmd.AddCommand(releaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gitlabCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gitlabCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
