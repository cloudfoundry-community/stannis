package agent

import (
	"github.com/cloudfoundry-community/gogobosh/models"
	"github.com/cloudfoundry-community/stannis/config"
)

// ToBOSH is the outbound data from a BOSH
type ToBOSH struct {
	Name        string             `form:"name"`
	TargetURI   string             `form:"target_uri"`
	UUID        string             `form:"uuid"`
	Version     string             `form:"version"`
	CPI         string             `form:"cpi"`
	Deployments models.Deployments `form:"deployments"`
}

// UploadGateway is the API for uploading deployments to Collector
type UploadGateway interface {
	UploadDeployments()
}

// Agent is the parent model for agent runtime behavior
type Agent struct {
	Config *config.AgentConfig
	Upload UploadGateway
}
