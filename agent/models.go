package agent

import (
	"github.com/cloudfoundry-community/gogobosh/models"
	"github.com/cloudfoundry-community/stannis/config"
)

// ToBOSH is the outbound data from a BOSH
type ToBOSH struct {
	Name        string             `json:"name"`
	Target      string             `json:"target"`
	UUID        string             `json:"uuid"`
	Version     string             `json:"version"`
	CPI         string             `json:"cpi"`
	Deployments models.Deployments `json:"deployments"`
}

// UploadGateway is the API for uploading deployments to Collector
type UploadGateway interface {
	UploadBulkDeployments()
	UploadDeploymentNames()
	UploadDeployments()
}

// Agent is the parent model for agent runtime behavior
type Agent struct {
	Config *config.AgentConfig
	Upload UploadGateway
}
