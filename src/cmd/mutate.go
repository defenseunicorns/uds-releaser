/*
Copyright © 2024 The Authors of uds-releaser

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

// mutateCmd represents the mutate command
var mutateCmd = &cobra.Command{
	Use:   "mutate [ flavor ]",
	Short: "Mutate version fields in the zarf.yaml and uds-bundle.yaml based on flavor",
	Args:  cobra.ExactArgs(1),
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

func init() {
	rootCmd.AddCommand(mutateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gitlabCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gitlabCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
