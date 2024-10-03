package utils

import (
	zarf "github.com/zarf-dev/zarf/src/api/v1alpha1"
)

func GetPackageName() (string, error) {
	var zarfPackage zarf.ZarfPackage
	err := LoadYaml("zarf.yaml", &zarfPackage)
	if err != nil {
		return "", err
	}

	return zarfPackage.Metadata.Name, nil
}
