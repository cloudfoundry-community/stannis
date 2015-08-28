package data

import (
	"sort"

	"github.com/cloudfoundry-community/stannis/upload"
)

// DeploymentsPerBOSH allows a BOSH's deployments to be indexed by BOSH UUID
type DeploymentsPerBOSH map[string]*BOSH

// BOSH describes a BOSH and its Deployments
type BOSH struct {
	Name        string
	Target      string
	ReallyUUID  string
	UUID        string
	Version     string
	CPI         string
	Deployments Deployments
}

// Deployments is a deployment-name to BOSHDeployment mapping
type Deployments map[string]*Deployment

// Deployment describes a BOSH deployment
type Deployment struct {
	Name     string
	Releases []struct {
		Name    string
		Version string
	}
	Stemcells []struct {
		Name    string
		Version string
	}
	CloudConfig string
}

// NewBOSH constructs a BOSH from the uploaded BOSH data
func NewBOSH(uploadedBOSH *upload.BOSH) (bosh *BOSH) {
	return &BOSH{
		Name:        uploadedBOSH.Name,
		Target:      uploadedBOSH.Target,
		ReallyUUID:  uploadedBOSH.ReallyUUID,
		UUID:        uploadedBOSH.UUID,
		Version:     uploadedBOSH.Version,
		CPI:         uploadedBOSH.CPI,
		Deployments: Deployments{},
	}
}

// UpdateDeployment adds/updates a Deployment from uploaded BOSHDeployment data
func (bosh *BOSH) UpdateDeployment(uploadedDeployment *upload.BOSHDeployment) {
	deployment := &Deployment{
		Name:        uploadedDeployment.Name,
		Releases:    uploadedDeployment.Releases,
		Stemcells:   uploadedDeployment.Stemcells,
		CloudConfig: uploadedDeployment.CloudConfig,
	}
	bosh.Deployments[deployment.Name] = deployment
}

// NewDeploymentsPerBOSH constructs a new mapping of Deployments to each BOSH
func NewDeploymentsPerBOSH() DeploymentsPerBOSH {
	return DeploymentsPerBOSH{}
}

// deploymentsPerRelease returns a {releaseName: []upload.BOSHDeployment} mapping
func (db DeploymentsPerBOSH) deploymentsPerRelease() (result map[string][]upload.BOSH) {
	result = map[string][]upload.BOSH{}
	for _, bosh := range db {
		for _, deployment := range bosh.Deployments {
			for _, release := range deployment.Releases {
				if result[release.Name] == nil {
					result[release.Name] = []upload.BOSH{}
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
