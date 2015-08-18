package data

import "github.com/cloudfoundry-community/bosh-pipeline-dashboard/upload"

// DeploymentsPerBOSH allows a BOSH's deployments to be indexed by BOSH UUID
type DeploymentsPerBOSH map[string]upload.UploadedFromBOSH

// NewDeploymentsPerBOSH constructs a new mapping of Deployments to each BOSH
func NewDeploymentsPerBOSH() DeploymentsPerBOSH {
	return DeploymentsPerBOSH{}
}
