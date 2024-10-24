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
