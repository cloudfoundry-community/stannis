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
		expectedDeployments *PipelinedDeployments
		db                  data.DeploymentsPerBOSH
		renderdata          *RenderData
	)
	BeforeEach(func() {
		expectedDeployments = TestScenarioData()
		db = data.NewDeploymentsPerBOSH()

		db.LoadFixtureData("fixtures/deployments-uuid-some-bosh-lite.json")
		db.LoadFixtureData("fixtures/deployments-uuid-aws-bosh-production.json")
		db.LoadFixtureData("fixtures/deployments-uuid-vsphere-bosh-sandbox.json")
	})

	Describe("Organize data based on pipeline configuration", func() {
		It("should have two tiers", func() {
			pipelineConfig, err := config.LoadConfigFromYAMLFile("../config/config.example.yml")
			Expect(err).NotTo(HaveOccurred())
			renderdata = NewRenderData(pipelineConfig)

			Expect(len(*expectedDeployments)).To(Equal(2))
		})
	})
})
