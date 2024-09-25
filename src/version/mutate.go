package version

import (
	"errors"
	"fmt"

	"github.com/defenseunicorns/uds-releaser/src/types"
	"github.com/defenseunicorns/uds-releaser/src/utils"
)

func MutateYamls(flavor types.Flavor) error {
	tagExists, err := utils.DoesTagExist("test")
	if err != nil {
		return err
	}

	if tagExists {
		fmt.Printf("Version %s exists in the git tags\n", flavor.Version)
		fmt.Print("No release necessary\n")

		return errors.New("version already exists")
	} else {
		fmt.Printf("Version %s does not exist in the git tags\n", flavor.Version)
		fmt.Print("Mutating package and bundle yamls\n")

		err = mutateZarfYaml(flavor)
		if err != nil {
			return err
		}

		err = mutateBundleYaml(flavor)
		if err != nil {
			return err
		}
	}
	return nil
}

func mutateZarfYaml(flavor types.Flavor) error {
	return nil
}

func mutateBundleYaml(flavor types.Flavor) error {
	return nil
}
