package rendertemplates

import (
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/config"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/data"
)

// PrepareRenderData constructs new renderdata based on pipeline config + latest BOSH deployments data
func PrepareRenderData(config *config.PipelinesConfig, data data.DeploymentsPerBOSH) *RenderData {
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

		for slotIndex := range configTier.Slots {
			deployments := Deployments{}
			slots[slotIndex] = &Slot{
				Deployments: deployments,
			}
		}
	}
	return renderdata
}
