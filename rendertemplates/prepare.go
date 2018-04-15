package rendertemplates

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/cloudfoundry-community/stannis/config"
	"github.com/cloudfoundry-community/stannis/data"
)

// PrepareRenderData constructs new renderdata based on pipeline config + latest BOSH deployments data
func PrepareRenderData(config *config.PipelinesConfig, db data.DeploymentsPerBOSH, filterByTag string) *RenderData {
	tiers := make([]*Tier, len(config.Tiers))
	filterTags := filterTagsForReleaseNames(db.ReleaseNames())
	renderdata := &RenderData{
		Config:     config,
		FilterTags: filterTags,
		Tiers:      tiers,
	}
	for tierIndex, configTier := range config.Tiers {
		slots := make([]*Slot, len(configTier.Slots))
		tiers[tierIndex] = &Tier{
			Name:  configTier.Name,
			Slots: slots,
		}

		for slotIndex, configSlot := range configTier.Slots {
			deployments := renderdata.DiscoverDeploymentsForSlot(db, configTier, configSlot, filterByTag)
			slots[slotIndex] = &Slot{
				Deployments: deployments,
			}
		}
	}
	return renderdata
}

// DiscoverDeploymentsForSlot searches through the database of known BOSH deployments for those
// that should appear in a configured tier/slot
func (renderdata *RenderData) DiscoverDeploymentsForSlot(db data.DeploymentsPerBOSH, configTier config.Tier, configSlot config.Slot, filterTag string) Deployments {
	var deployments Deployments
	for _, boshDeployments := range db {
		fmt.Printf("BOSH deployments: %#v\n", boshDeployments)
		keys := make([]string, 0)
		for key := range boshDeployments.Deployments {
			keys = append(keys, key)
		}

		sort.Strings(keys)
		for _, key := range keys {
			boshDeployment := boshDeployments.Deployments[key]
			match := false
			if configSlot.Filter.DeploymentNameRegexp != "" {
				match, _ = regexp.MatchString(configSlot.Filter.DeploymentNameRegexp, boshDeployment.Name)
				if match {
					deployment := NewDeployment(configTier, configSlot, boshDeployment)
					if deployment.ContainsFilterTag(filterTag) {
						deployments = append(deployments, deployment)
					}
				}
			}
			// TODO: also allow filter via TargetURI
			if !match && configSlot.Filter.BoshUUID != "" {
				if boshDeployments.UUID == configSlot.Filter.BoshUUID {
					match = true
					deployment := NewDeployment(configTier, configSlot, boshDeployment)
					if deployment.ContainsFilterTag(filterTag) {
						deployments = append(deployments, deployment)
					}
				}
			}
			if !match && configSlot.Filter.TargetName != "" {
				if boshDeployments.Name == configSlot.Filter.TargetName {
					match = true
					deployment := NewDeployment(configTier, configSlot, boshDeployment)
					if deployment.ContainsFilterTag(filterTag) {
						deployments = append(deployments, deployment)
					}
				}
			}
			if !match && configSlot.Filter.TargetURI != "" {
				if boshDeployments.Target == configSlot.Filter.TargetURI {
					match = true
					deployment := NewDeployment(configTier, configSlot, boshDeployment)
					if deployment.ContainsFilterTag(filterTag) {
						deployments = append(deployments, deployment)
					}
				}
			}
		}
	}
	return deployments
}

func filterTagsForReleaseNames(releaseNames []string) (tags []FilterTag) {
	tags = make([]FilterTag, len(releaseNames))
	for i, name := range releaseNames {
		tags[i] = FilterTag{name, name, "icon-cloud"}
	}
	return
}
