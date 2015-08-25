package agent_test

import (
	. "github.com/cloudfoundry-community/stannis/agent"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Convert bosh target/uuid into ReallyUUID", func() {
	It("BOSH Target just IP", func() {
		Expect(ReallyUUID("1.2.3.4", "some-uuid")).To(Equal("1.2.3.4-some-uuid"))
	})
	It("BOSH Target https://IP:25555", func() {
		Expect(ReallyUUID("https://1.2.3.4:25555", "some-uuid")).To(Equal("1.2.3.4-some-uuid"))
	})
	It("BOSH Target just hostname", func() {
		Expect(ReallyUUID("my.bosh.com", "some-uuid")).To(Equal("my.bosh.com-some-uuid"))
	})
	It("BOSH Target https://hostname:25555", func() {
		Expect(ReallyUUID("https://my.bosh.com:25555", "some-uuid")).To(Equal("my.bosh.com-some-uuid"))
	})
})
