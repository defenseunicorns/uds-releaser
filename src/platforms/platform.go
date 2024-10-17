package platforms

import (
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
