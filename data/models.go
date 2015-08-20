package data

import (
	"encoding/json"
	"io/ioutil"

	"github.com/cloudfoundry-community/stannis/upload"
)

// DeploymentsPerBOSH allows a BOSH's deployments to be indexed by BOSH UUID
type DeploymentsPerBOSH map[string]upload.FromBOSH

// NewDeploymentsPerBOSH constructs a new mapping of Deployments to each BOSH
func NewDeploymentsPerBOSH() DeploymentsPerBOSH {
	return DeploymentsPerBOSH{}
}

// LoadFixtureData is a text helper
func (db DeploymentsPerBOSH) LoadFixtureData(path string) (err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	deployments := &upload.FromBOSH{}
	err = json.Unmarshal(bytes, &deployments)

	db[deployments.UUID] = *deployments
	return
}
