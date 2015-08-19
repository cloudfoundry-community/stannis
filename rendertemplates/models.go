package rendertemplates

import (
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/config"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/data"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/upload"
)

// The PipelinedDeployments struct is used by the dashboard template to render/display
// the BOSH deployments.
// The BOSH deployments are filtered - say only displaying
// Cloud Foundry deployments (those using 'cf' release), but no others.
// The BOSH deployments are laid out on the dashboard - and hence structured in
// PipelinedDeployments - based on the config.PipelinesConfig.
//
// The dashboard is laid out in "tiers" - rows down the page - that might represent
// a group of deployments that are related - such as deployments in a common datacenter.
// Within each "tier", the deployments are grouped into "columns" (could be thought
// of as "slots"). Typically there might be a single deployment in each column -
// typically representing a deployment in a sequenced pipeline of deployments,
// sandbox -> pre-production -> production.
//
// To display a deployment is to display its name and the set of BOSH releases
// and BOSH stemcells being used, and the versions of them. This is the primary
// purpose of the dashboard view: what/where is software/versions running?
//
// It is assumed that deployments in the far right columns should be using BOSH release
// versions that are older (smaller version number) than the deployments in columns
// to the left. As such, visual indicators are given to show that a BOSH release
// version is higher or lower than the deployment to its immediate left.

// RenderData is a collection of deployments in the Director by tiers/pipelines
type RenderData struct {
	Config      *config.PipelinesConfig
	Deployments PipelinedDeployments
}

// PipelinedDeployments is a collection of deployments in the Director by tiers/pipelines
type PipelinedDeployments []*Tier

// Tier is a collection of deployments in the Director
type Tier []*Slot

// Slot in the dashboard that displays some deployments
type Slot []*Deployment

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
		&Tier{
			&Slot{
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
				&Deployment{
					Name: "try-anything / bosh-lite-2 - cf2-try-anything",
					Releases: []NameVersion{
						NameVersion{Name: "cf", Version: "214", DisplayClass: "icon-arrow-up green"},
						NameVersion{Name: "cf-sensu-client", Version: "2", DisplayClass: "icon-arrow-up green"},
					},
					Stemcells: []NameVersion{
						NameVersion{Name: "warden", Version: "2776", DisplayClass: "icon-minus blue"},
					},
				},
			},
		},
		&Tier{
			&Slot{
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
			},
			&Slot{
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
			},
			&Slot{
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
		},
	}
}
