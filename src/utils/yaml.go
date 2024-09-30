package utils

import (
	"os"

	"github.com/defenseunicorns/uds-releaser/src/types"
	goyaml "github.com/goccy/go-yaml"
)

func LoadReleaserConfig() (types.ReleaserConfig, error) {

	var config types.ReleaserConfig
	err := LoadYaml("releaser.yaml", &config)
	if err != nil {
		return types.ReleaserConfig{}, err
	}

	return config, nil
}

func LoadYaml(path string, destVar interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = goyaml.Unmarshal(data, destVar)
	if err != nil {
		return err
	}
	return nil
}

func UpdateYaml(path string, srcVar interface{}) error {
	data, err := goyaml.Marshal(srcVar)
	if err != nil {
		return err
	}

	yamlInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, yamlInfo.Mode())
	if err != nil {
		return err
	}
	return nil
}
