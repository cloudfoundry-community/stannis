package rendertemplates_test

import (
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/config"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/data"
	. "github.com/cloudfoundry-community/bosh-pipeline-dashboard/rendertemplates"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Prepare data for templates", func() {
	var (
		pipelineConfig *config.PipelinesConfig
		expectedTiers  *Tiers
		db             data.DeploymentsPerBOSH
		renderdata     *RenderData
	)
	BeforeEach(func() {
		expectedTiers = TestScenarioData()
		db = data.NewDeploymentsPerBOSH()

		db.LoadFixtureData("fixtures/deployments-uuid-some-bosh-lite.json")
		db.LoadFixtureData("fixtures/deployments-uuid-aws-bosh-production.json")
		db.LoadFixtureData("fixtures/deployments-uuid-vsphere-bosh-sandbox.json")

		var err error
		pipelineConfig, err = config.LoadConfigFromYAMLFile("../config/config.example.yml")
		Expect(err).NotTo(HaveOccurred())

		renderdata = PrepareRenderData(pipelineConfig, db)
	})

	It("has tiers", func() {
		Expect(len(renderdata.Tiers)).To(Equal(len(*expectedTiers)))
	})
})
