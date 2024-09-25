package version

import (
	"errors"
	"fmt"

	"github.com/defenseunicorns/uds-releaser/src/types"
	"github.com/defenseunicorns/uds-releaser/src/utils"
	zarf "github.com/zarf-dev/zarf/src/api/v1alpha1"
	uds "github.com/defenseunicorns/uds-cli/src/types"
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

		packageName, err := mutateZarfYaml(flavor)
		if err != nil {
			return err
		}

		err = mutateBundleYaml(flavor, packageName)
		if err != nil {
			return err
		}
	}
	return nil
}

func mutateZarfYaml(flavor types.Flavor) (packageName string, err error) {
	var zarfPackage zarf.ZarfPackage
	err = utils.LoadYaml("zarf.yaml", &zarfPackage)
	if err != nil {
		return "", err
	}

	zarfPackage.Metadata.Version = flavor.Version

	err = utils.UpdateYaml("zarf.yaml", zarfPackage)
	if err != nil {
		return zarfPackage.Metadata.Name, err
	}

	return zarfPackage.Metadata.Name, nil
}

func mutateBundleYaml(flavor types.Flavor, packageName string) error {
	var bundle uds.UDSBundle
	err := utils.LoadYaml("bundle/uds-bundle.yaml", &bundle)
	if err != nil {
		return err
	}

	bundle.Metadata.Version = flavor.Version

	// Find the package that matches the package name and update its ref
	for i, bundledPackage := range bundle.Packages {
		if bundledPackage.Name == packageName {
			bundle.Packages[i].Ref = flavor.Version
		}
	}

	err = utils.UpdateYaml("bundle/uds-bundle.yaml", bundle)
	if err != nil {
		return err
	}
	return nil
}
