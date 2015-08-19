package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// PipelinesConfig describes how pipelines will be displayed, how to allocate discovered deployments
type PipelinesConfig struct {
	Pipelines []struct {
		Name   string `yaml:"name"`
		Tag    string `yaml:"tag"`
		Filter struct {
			UsingBoshReleases []string `yaml:"using_bosh_releases"`
		} `yaml:"filter"`
	} `yaml:"pipelines"`
	Tiers []struct {
		Name  string `yaml:"name"`
		Slots []struct {
			Name   string `yaml:"name"`
			Filter struct {
				BoshUUID             string `yaml:"bosh_uuid"`
				DeploymentNameRegexp string `yaml:"deployment_name_regexp"`
			} `yaml:"filter"`
		} `yaml:"slots"`
	} `yaml:"tiers"`
}

// LoadConfigFromYAMLFile loads pipeline configuration from a YAML file
func LoadConfigFromYAMLFile(path string) (config *PipelinesConfig, err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &PipelinesConfig{}
	err = yaml.Unmarshal(bytes, &config)

	return config, err
}
