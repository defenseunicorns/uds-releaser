package utils

import (
	"errors"

	"github.com/defenseunicorns/uds-releaser/src/types"
)

func GetFlavorConfig(flavor string, config types.ReleaserConfig) (types.Flavor, error) {
	for _, f := range config.Flavors {
		if f.Name == flavor {
			return f, nil
		}
	}
	return types.Flavor{}, errors.New("flavor not found")
}