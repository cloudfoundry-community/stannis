package agent_test

import (
	. "github.com/cloudfoundry-community/stannis/agent"
	"github.com/cloudfoundry-community/stannis/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Prepare data for templates", func() {
	var (
		agentConfig *config.AgentConfig
		subject     Agent
	)
	BeforeEach(func() {
		var err error
		agentConfig, err = config.LoadAgentConfigFromYAMLFile("../config/agent.config.example.yml")
		Expect(err).NotTo(HaveOccurred())

		fakeUploadGateway := NewFakeUploadGateway()

		subject = Agent{
			Config: agentConfig,
			Upload: fakeUploadGateway,
		}
	})

	It("has max bulk upload", func() {
		Expect(agentConfig.MaxBulkUploadSize).To(Equal(5))
	})

	// Due to the 100-continue issue, we upload only small data sets
	// https://github.com/golang/go/issues/3665
	It("uploads small sets of deployments in bulk", func() {
	})

	It("uploads only names of deployments first then uploads each deployment", func() {
	})
})

type FakeUploadGateway struct{}

func NewFakeUploadGateway() (gateway FakeUploadGateway) {
	return
}

func (gateway FakeUploadGateway) UploadBulkDeployments() {}
func (gateway FakeUploadGateway) UploadDeploymentNames() {}
func (gateway FakeUploadGateway) UploadDeployments()     {}
