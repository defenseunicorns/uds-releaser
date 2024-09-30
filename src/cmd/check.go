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
	"errors"
	"fmt"

	"github.com/defenseunicorns/uds-releaser/src/utils"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check [ flavor ]",
	Short: "check if release is necessary for given flavor",
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

		rootCmd.SilenceUsage = true

		versionAndFlavor := fmt.Sprintf("%s-%s", currentFlavor.Version, currentFlavor.Name)

		tagExists, err := utils.DoesTagExist(versionAndFlavor)
		if err != nil {
			return err
		}
		if tagExists {
			fmt.Printf("Version %s is already tagged\n", versionAndFlavor)
			return errors.New("no release necessary")
		} else {
			fmt.Printf("Version %s is not tagged\n", versionAndFlavor)
			return nil
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gitlabCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gitlabCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
