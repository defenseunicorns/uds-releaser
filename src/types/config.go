package types

type Flavor struct {
	Name              string `yaml:"name"`
	Version           string `yaml:"version"`
	PublishBundle     bool   `yaml:"publishBundle,omitempty,default=false"`
	PublishPackageUrl string `yaml:"publishPackageUrl"`
	PublishBundleUrl  string `yaml:"publishBundleUrl"`
}

type ReleaserConfig struct {
	Flavors []Flavor `yaml:"flavors"`
}