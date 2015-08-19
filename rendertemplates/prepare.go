package rendertemplates

import (
	"regexp"

	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/config"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/data"
)

// PrepareRenderData constructs new renderdata based on pipeline config + latest BOSH deployments data
func PrepareRenderData(config *config.PipelinesConfig, db data.DeploymentsPerBOSH) *RenderData {
	tiers := make([]*Tier, len(config.Tiers))
	renderdata := &RenderData{
		Config: config,
		Tiers:  tiers,
	}
	for tierIndex, configTier := range config.Tiers {
		slots := make([]*Slot, len(configTier.Slots))
		tiers[tierIndex] = &Tier{
			Name:  configTier.Name,
			Slots: slots,
		}

		for slotIndex, configSlot := range configTier.Slots {
			deployments := renderdata.DiscoverDeploymentsForSlot(db, configTier, configSlot)
			slots[slotIndex] = &Slot{
				Deployments: deployments,
			}
		}
	}
	return renderdata
}

// DiscoverDeploymentsForSlot searches through the database of known BOSH deployments for those
// that should appear in a configured tier/slot
func (renderdata *RenderData) DiscoverDeploymentsForSlot(db data.DeploymentsPerBOSH, configTier config.Tier, configSlot config.Slot) Deployments {
	var deployments Deployments
	for _, boshDeployments := range db {
		for _, boshDeployment := range boshDeployments.Deployments {
			match := false
			if configSlot.Filter.DeploymentNameRegexp != "" {
				match, _ = regexp.MatchString(configSlot.Filter.DeploymentNameRegexp, boshDeployment.Name)
				if match {
					deployments = append(deployments, NewDeployment(boshDeployment))
				}
			}
			if !match && configSlot.Filter.BoshUUID != "" {
				if boshDeployments.UUID == configSlot.Filter.BoshUUID {
					deployments = append(deployments, NewDeployment(boshDeployment))
				}
			}
		}
	}
	return deployments
}
