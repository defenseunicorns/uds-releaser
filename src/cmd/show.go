// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package cmd

import (
	"fmt"

	"github.com/defenseunicorns/uds-releaser/src/utils"
	"github.com/spf13/cobra"
)

var showVersionOnly bool

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show flavor",
	Short: "Show the current version for a given flavor",
	Args:  cobra.ExactArgs(1),
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

		if showVersionOnly {
			fmt.Printf("%s\n", currentFlavor.Version)
		} else {
			fmt.Printf("%s-%s\n", currentFlavor.Version, currentFlavor.Name)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().BoolVarP(&showVersionOnly, "version-only", "v", false, "Show only the version without flavor appended")
}
