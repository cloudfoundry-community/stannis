package upload

// FromBOSH is the inbound data from a BOSH
type FromBOSH struct {
	Name        string                `form:"name"`
	Target      string                `form:"target"`
	UUID        string                `form:"uuid"`
	Version     string                `form:"version"`
	CPI         string                `form:"cpi"`
	Deployments []*DeploymentFromBOSH `form:"deployments"`
}

// DeploymentFromBOSH is the received list of deployments from a BOSH
type DeploymentFromBOSH struct {
	Name     string `form:"name"`
	Releases []struct {
		Name    string `form:"name"`
		Version string `form:"version"`
	} `form:"releases"`
	Stemcells []struct {
		Name    string `form:"name"`
		Version string `form:"version"`
	} `form:"stemcells"`
	CloudConfig string `form:"cloud_config"`
}
