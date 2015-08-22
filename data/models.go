package data

import (
	"encoding/json"
	"io/ioutil"
	"sort"

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

// deploymentsPerRelease returns a {releaseName: []upload.DeploymentFromBOSH} mapping
func (db DeploymentsPerBOSH) deploymentsPerRelease() (result map[string][]upload.FromBOSH) {
	result = map[string][]upload.FromBOSH{}
	for _, bosh := range db {
		for _, deployment := range bosh.Deployments {
			for _, release := range deployment.Releases {
				if result[release.Name] == nil {
					result[release.Name] = []upload.FromBOSH{}
				}
			}
		}
	}
	return
}

// ReleaseNames returns the names of the BOSH releases used by deployments
func (db DeploymentsPerBOSH) ReleaseNames() (names []string) {
	deploymentsPerRelease := db.deploymentsPerRelease()
	for release := range deploymentsPerRelease {
		names = append(names, release)
	}
	sort.Strings(names)
	return
}
