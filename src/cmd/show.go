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
