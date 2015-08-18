package rendertemplates_test

import (
	. "github.com/cloudfoundry-community/bosh-pipeline-dashboard/rendertemplates"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Prepare data for templates", func() {
	var (
		expectedDeployments *PipelinedDeployments
	)
	BeforeEach(func() {
		expectedDeployments = ExampleData()
	})

	It("should have two tiers", func() {
		Expect(len(*expectedDeployments)).To(Equal(2))
	})
})
