// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package platforms

import (
	"fmt"
	"os"

	"github.com/defenseunicorns/uds-releaser/src/types"
	"github.com/defenseunicorns/uds-releaser/src/utils"
)

type Platform interface {
	TagAndRelease(flavor types.Flavor, tokenVarName string) error
}

func LoadAndTag(releaserDir, flavor, tokenVarName string, platform Platform) error {
	err := VerifyEnvVar(tokenVarName)
	if err != nil {
		return err
	}

	releaserConfig, err := utils.LoadReleaserConfig(releaserDir)
	if err != nil {
		return err
	}

	currentFlavor, err := utils.GetFlavorConfig(flavor, releaserConfig)
	if err != nil {
		return err
	}

	return platform.TagAndRelease(currentFlavor, tokenVarName)
}

func VerifyEnvVar(varName string) error {
	if value, exists := os.LookupEnv(varName); !exists || value == "" {
		return fmt.Errorf("%s is unset or empty", varName)
	}

	return nil
}
