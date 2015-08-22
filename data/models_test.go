package data_test

import (
	. "github.com/cloudfoundry-community/stannis/data"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Data", func() {
	var (
		db DeploymentsPerBOSH
	)
	BeforeEach(func() {
		db = NewDeploymentsPerBOSH()
		db.LoadFixtureData("../rendertemplates/fixtures/deployments-uuid-some-bosh-lite.json")
		db.LoadFixtureData("../rendertemplates/fixtures/deployments-uuid-aws-bosh-production.json")
		db.LoadFixtureData("../rendertemplates/fixtures/deployments-uuid-vsphere-bosh-sandbox.json")
	})

	It("finds releases", func() {
		Expect(db.ReleaseNames()).To(Equal([]string{"cf", "cf-haproxy", "concourse", "garden-linux"}))
	})
})
