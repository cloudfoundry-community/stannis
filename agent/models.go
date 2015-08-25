package agent

import "github.com/cloudfoundry-community/gogobosh/models"

// ToBOSH is the outbound data from a BOSH
type ToBOSH struct {
	Name        string             `form:"name"`
	TargetURI   string             `form:"target_uri"`
	UUID        string             `form:"uuid"`
	Version     string             `form:"version"`
	CPI         string             `form:"cpi"`
	Deployments models.Deployments `form:"deployments"`
}
