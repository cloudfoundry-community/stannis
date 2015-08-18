package rendertemplates_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBOSHDashboard(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BOSH Dashboard suite")
}
