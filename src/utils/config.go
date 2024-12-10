// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package utils

import (
	"errors"

	"github.com/defenseunicorns/uds-pk/src/types"
)

func GetFlavorConfig(flavor string, config types.ReleaseConfig) (types.Flavor, error) {
	for _, f := range config.Flavors {
		if f.Name == flavor {
			return f, nil
		}
	}
	return types.Flavor{}, errors.New("flavor not found")
}
