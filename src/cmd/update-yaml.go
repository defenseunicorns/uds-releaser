// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package cmd

import (
	"github.com/defenseunicorns/uds-releaser/src/utils"
	"github.com/defenseunicorns/uds-releaser/src/version"
	"github.com/spf13/cobra"
)

// updateYamlCmd represents the updateyaml command
var updateYamlCmd = &cobra.Command{
	Use:     "update-yaml flavor",
	Aliases: []string{"u"},
	Short:   "Update the version fields in the zarf.yaml and uds-bundle.yaml based on flavor",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		releaserConfig, err := utils.LoadReleaserConfig(releaserDir)
		if err != nil {
			return err
		}

		currentFlavor, err := utils.GetFlavorConfig(args[0], releaserConfig)
		if err != nil {
			return err
		}

		rootCmd.SilenceUsage = true

		return version.UpdateYamls(currentFlavor)
	},
}

func init() {
	rootCmd.AddCommand(updateYamlCmd)
}
