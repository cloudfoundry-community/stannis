package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// AgentConfig configures the `agent` subcommand to send information back to
// the webserver collector
type AgentConfig struct {
	BOSHTarget   string `yaml:"bosh_target"`
	BOSHUsername string `yaml:"bosh_username"`
	BOSHPassword string `yaml:"bosh_password"`

	WebserverTarget   string `yaml:"webserver_target"`
	WebserverUsername string `yaml:"webserver_username"`
	WebserverPassword string `yaml:"webserver_password"`

	MaxBulkUploadSize int `yaml:"max_bulk_upload_size"` // defaults to 5 below
}

// LoadAgentConfigFromYAMLFile loads pipeline configuration from a YAML file
func LoadAgentConfigFromYAMLFile(path string) (config *AgentConfig, err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &AgentConfig{}
	err = yaml.Unmarshal(bytes, &config)

	config.MaxBulkUploadSize = 0

	return config, err
}
