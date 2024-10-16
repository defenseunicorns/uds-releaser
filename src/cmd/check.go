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
	"errors"
	"fmt"

	"github.com/defenseunicorns/uds-releaser/src/utils"
	"github.com/spf13/cobra"
)

var checkBoolOutput bool

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check flavor",
	Short: "Check if release is necessary for given flavor",
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

		versionAndFlavor := fmt.Sprintf("%s-%s", currentFlavor.Version, currentFlavor.Name)

		tagExists, err := utils.DoesTagExist(versionAndFlavor)
		if err != nil {
			return err
		}
		if tagExists {
			if checkBoolOutput {
				fmt.Println("false")
			} else {
				fmt.Printf("Version %s is already tagged\n", versionAndFlavor)
				return errors.New("no release necessary")
			}
		} else {
			if checkBoolOutput {
				fmt.Println("true")
			} else {
				fmt.Printf("Version %s is not tagged\n", versionAndFlavor)
			}
		}
		return nil
	},
}

func init() {
	checkCmd.Flags().BoolVarP(&checkBoolOutput, "boolean", "b", false, "Switch the output string to a true/false based on if a release is necessary. True if a release is necessary, false if not.")
	rootCmd.AddCommand(checkCmd)
}
