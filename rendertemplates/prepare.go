package rendertemplates

import (
	"fmt"

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
func (renderdata *RenderData) DiscoverDeploymentsForSlot(db data.DeploymentsPerBOSH, configTier config.Tier, configSlot config.Slot) (deployments Deployments) {
	fmt.Println(configTier)
	fmt.Println(configSlot)
	for _, boshDeployments := range db {
		boshName := boshDeployments.Name
		deployments := boshDeployments.Deployments
		fmt.Println(boshName)
		fmt.Println(deployments)
	}
	return
}
