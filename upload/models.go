package upload

// UploadedFromBOSH is the inbound data from a BOSH
type UploadedFromBOSH struct {
	UUID        string                       `form:"uuid"`
	Name        string                       `form:"name"`
	Deployments []UploadedDeploymentFromBOSH `form:"deployments"`
}

// UploadedDeploymentFromBOSH is the received list of deployments from a BOSH
type UploadedDeploymentFromBOSH struct {
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
