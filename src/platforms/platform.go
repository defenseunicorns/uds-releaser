// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package platforms

import (
	"regexp"

	"github.com/defenseunicorns/uds-releaser/src/types"
	"github.com/defenseunicorns/uds-releaser/src/utils"
)

type Platform interface {
	TagAndRelease(flavor types.Flavor, tokenVarName string) error
}

func LoadAndTag(releaserDir, flavor, tokenVarName string, platform Platform) error {
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

func ReleaseExists(expectedStatusCode, receivedStatusCode int, response string, pattern string) bool {
	return receivedStatusCode == expectedStatusCode && regexp.MustCompile(pattern).MatchString(response)
}
