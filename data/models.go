package data

import (
	"fmt"
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
	ExtraData   ExtraData
}

// ExtraData is uploaded data about a running Deployment
type ExtraData map[string]*DeploymentData

// DeploymentData describes extra data about a running Deployment
type DeploymentData struct {
	ReallyUUID     string
	DeploymentName string
	Label          string
	Data           []struct {
		Indicator string
		Value     string
		Label     string
	}
}

// UpdateBOSH constructs a BOSH from the uploaded BOSH data
func (db DeploymentsPerBOSH) UpdateBOSH(uploadedBOSH *upload.BOSH) {
	reallyUUID := uploadedBOSH.ReallyUUID
	if db[reallyUUID] == nil {
		bosh := BOSH{
			Deployments: Deployments{},
		}
		db[reallyUUID] = &bosh
	}
	bosh := db[reallyUUID]
	bosh.Name = uploadedBOSH.Name
	bosh.Target = uploadedBOSH.Target
	bosh.ReallyUUID = uploadedBOSH.ReallyUUID
	bosh.UUID = uploadedBOSH.UUID
	bosh.Version = uploadedBOSH.Version
	bosh.CPI = uploadedBOSH.CPI

	fmt.Println(uploadedBOSH.ReallyUUID)
	fmt.Println(db)
}

// UpdateDeployment adds/updates a Deployment from uploaded BOSHDeployment data
func (bosh *BOSH) UpdateDeployment(uploadedDeployment *upload.BOSHDeployment) {
	name := uploadedDeployment.Name
	if bosh.Deployments[name] == nil {
		deployment := &Deployment{
			ExtraData: ExtraData{},
		}
		bosh.Deployments[name] = deployment
	}
	deployment := bosh.Deployments[name]
	deployment.Name = uploadedDeployment.Name
	deployment.Releases = uploadedDeployment.Releases
	deployment.Stemcells = uploadedDeployment.Stemcells
	deployment.CloudConfig = uploadedDeployment.CloudConfig
}

// UpdateDeploymentData adds/updates addition data about a BOSH deployment in action
func (deployment *Deployment) UpdateDeploymentData(uploadedData *upload.DeploymentData) {
	data := &DeploymentData{
		ReallyUUID:     uploadedData.ReallyUUID,
		DeploymentName: uploadedData.DeploymentName,
		Label:          uploadedData.Label,
		Data:           uploadedData.Data,
	}
	deployment.ExtraData[data.Label] = data
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
