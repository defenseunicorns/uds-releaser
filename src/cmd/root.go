// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var releaserDir string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "uds-releaser",
	Short: "UDS Releaser is a tool for releasing UDS packages",
	Long: `UDS Releaser is a tool that facilitates the release
	of UDS packages. It provides commands for checking if a release is necessary,
	mutating version fields in the zarf.yaml and uds-bundle.yaml files, and creating tags
	and releases.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&releaserDir, "releaser-dir", "d", ".", "Path to the directory containing the releaser.yaml file")
}
