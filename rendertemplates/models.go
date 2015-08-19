package rendertemplates

import (
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/config"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/data"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/upload"
)

//Transposing the BOSH deployments data and mapping them into columns
// For each

// RenderData is a collection of deployments in the Director by tiers/pipelines
type RenderData struct {
	Config      *config.PipelinesConfig
	Deployments PipelinedDeployments
}

// PipelinedDeployments is a collection of deployments in the Director by tiers/pipelines
type PipelinedDeployments []*Deployments

// Deployments is a collection of deployments in the Director
type Deployments []*Deployment

// Deployment describes a running BOSH deployment and the
// Releases and Stemcells it is using.
type Deployment struct {
	Name      string
	Releases  []NameVersion
	Stemcells []NameVersion
}

// NameVersion is a reusable structure for Name/Version information
type NameVersion struct {
	Name         string
	Version      string
	DisplayClass string
}

// NewRenderData constructs new renderdata
func NewRenderData(config *config.PipelinesConfig) *RenderData {
	return &RenderData{Config: config}
}

// PrepareDeployments converts data into structures used by dashboard template
func (renderdata *RenderData) PrepareDeployments(data data.DeploymentsPerBOSH) {
	// TODO: structure the output based on pipeline configuration
	for _, boshDeployments := range data {
		renderdata.addBOSHDeployments(boshDeployments)
	}
}

func (renderdata RenderData) addBOSHDeployments(data upload.UploadedFromBOSH) {
	// deployments := &Deployments{}
}

// TestScenarioData returns some example data
func TestScenarioData() *PipelinedDeployments {
	return &PipelinedDeployments{

		&Deployments{
			&Deployment{
				Name: "try-anything / bosh-lite - cf-try-anything",
				Releases: []NameVersion{
					NameVersion{Name: "cf", Version: "214", DisplayClass: "icon-arrow-up green"},
					NameVersion{Name: "cf-sensu-client", Version: "1", DisplayClass: "icon-minus blue"},
				},
				Stemcells: []NameVersion{
					NameVersion{Name: "warden", Version: "2776", DisplayClass: "icon-minus blue"},
				},
			},
		},
		&Deployments{
			&Deployment{
				Name: "legacy / sandbox / aws - cf-sandbox-r5",
				Releases: []NameVersion{
					NameVersion{Name: "cf", Version: "211", DisplayClass: "icon-arrow-down red"},
					NameVersion{Name: "cf-sensu-client", Version: "1", DisplayClass: "icon-minus blue"},
				},
				Stemcells: []NameVersion{
					NameVersion{Name: "aws", Version: "3033", DisplayClass: "icon-minus blue"},
				},
			},
			&Deployment{
				Name: "legacy / dev / aws - cf-devprod-r2",
				Releases: []NameVersion{
					NameVersion{Name: "cf", Version: "211", DisplayClass: "icon-minus blue"},
					NameVersion{Name: "cf-sensu-client", Version: "1", DisplayClass: "icon-minus blue"},
				},
				Stemcells: []NameVersion{
					NameVersion{Name: "aws", Version: "3033", DisplayClass: "icon-minus blue"},
				},
			},
			&Deployment{
				Name: "legacy / prod / aws - prod-cloudfoundry",
				Releases: []NameVersion{
					NameVersion{Name: "cf", Version: "205", DisplayClass: "icon-arrow-down red"},
					NameVersion{Name: "cf-sensu-client", Version: "1", DisplayClass: "icon-minus blue"},
				},
				Stemcells: []NameVersion{
					NameVersion{Name: "aws", Version: "3000", DisplayClass: "icon-arrow-down red"},
				},
			},
		},
	}
}
