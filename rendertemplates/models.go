package rendertemplates

import (
	"fmt"
	"regexp"

	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/config"
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
	Config *config.PipelinesConfig
	Tiers  Tiers
}

// Tiers is a collection of deployments in the Director by tiers/pipelines
type Tiers []*Tier

// Tier is a collection of deployments in the Director
type Tier struct {
	Name  string
	Slots Slots
}

// Slots is an ordered sequence of slots in a pipeline of deployments within a Tier
type Slots []*Slot

// Slot in the dashboard that displays some deployments
type Slot struct {
	Deployments Deployments
}

// Deployments is a set of deployments
type Deployments []*Deployment

// Deployment describes a running BOSH deployment and the
// Releases and Stemcells it is using.
type Deployment struct {
	Name      string
	Releases  []DisplayNameVersion
	Stemcells []DisplayNameVersion
}

// DisplayNameVersion is a reusable structure for Name/Version information
type DisplayNameVersion struct {
	Name         string
	Version      string
	DisplayClass string
}

// NewDeployment converts BOSH deployment information into a deployment view for the dashboard
func NewDeployment(configTier config.Tier, configSlot config.Slot, boshDeployment upload.DeploymentFromBOSH) (deployment *Deployment) {
	tierName := configTier.Name
	slotName := configSlot.Name
	cpi := "no-stemcell"
	if len(boshDeployment.Stemcells) > 0 {
		cpi = cpiFromStemcell(boshDeployment.Stemcells[0].Name)
	}

	name := fmt.Sprintf("%s / %s / %s - %s", tierName, slotName, cpi, boshDeployment.Name)
	if slotName == cpi {
		name = fmt.Sprintf("%s / %s - %s", tierName, slotName, boshDeployment.Name)
	}

	releases := make([]DisplayNameVersion, len(boshDeployment.Releases))
	for releaseIndex := range releases {
		boshRelease := boshDeployment.Releases[releaseIndex]
		releases[releaseIndex] = DisplayNameVersion{
			Name:         boshRelease.Name,
			Version:      boshRelease.Version,
			DisplayClass: "icon-minus blue",
		}
	}

	stemcells := make([]DisplayNameVersion, len(boshDeployment.Stemcells))
	for stemcellIndex := range stemcells {
		boshStemcell := boshDeployment.Stemcells[stemcellIndex]
		stemcells[stemcellIndex] = DisplayNameVersion{
			Name:         boshStemcell.Name,
			Version:      boshStemcell.Version,
			DisplayClass: "icon-minus blue",
		}
	}

	deployment = &Deployment{
		Name:      name,
		Releases:  releases,
		Stemcells: stemcells,
	}
	return
}

func cpiFromStemcell(stemcellName string) string {
	match, _ := regexp.MatchString("boshlite", stemcellName)
	if match {
		return "bosh-lite"
	}
	match, _ = regexp.MatchString("vsphere", stemcellName)
	if match {
		return "vsphere"
	}
	match, _ = regexp.MatchString("aws", stemcellName)
	if match {
		return "aws"
	}
	match, _ = regexp.MatchString("openstack", stemcellName)
	if match {
		return "openstack"
	}
	match, _ = regexp.MatchString("vcloud", stemcellName)
	if match {
		return "vcloud"
	}
	return "unknown-stemcell"
}
