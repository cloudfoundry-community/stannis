package upload

// BOSH is the inbound data from a BOSH
type BOSH struct {
	Name       string `form:"name"`
	Target     string `form:"target"`
	ReallyUUID string `form:"reallyuuid"`
	UUID       string `form:"uuid"`
	Version    string `form:"version"`
	CPI        string `form:"cpi"`
}

// BOSHDeployment is the received list of deployments from a BOSH
type BOSHDeployment struct {
	ReallyUUID string `form:"reallyuuid"`
	Name       string `form:"name"`
	Releases   []struct {
		Name    string
		Version string
	} `form:"releases"`
	Stemcells []struct {
		Name    string
		Version string
	} `form:"stemcells"`
	CloudConfig string `form:"cloudconfig"`
}

// ExtraData captures some extra data about a deployment from a plugin
type ExtraData struct {
}
