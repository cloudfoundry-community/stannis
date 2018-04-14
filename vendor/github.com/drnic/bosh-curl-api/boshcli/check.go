package boshcli

import (
	"fmt"
	"log"
	"os/exec"
)

// Check that `bosh curl` exists else error & exit
func Check() {
	// check that 'bosh' available
	cmd := exec.Command("sh", "-c", "bosh -h")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", stdoutStderr)
		log.Fatal("Install 'bosh' from https://github.com/cloudfoundry/bosh-cli/pull/408")
	}

	// check that 'bosh curl' available
	cmd = exec.Command("sh", "-c", "bosh curl -h")
	stdoutStderr, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", stdoutStderr)
		log.Fatal("Need 'bosh curl' from https://github.com/cloudfoundry/bosh-cli/pull/408")
	}

	// check that bosh environment configured and connectable
	cmd = exec.Command("sh", "-c", "bosh curl /info")
	stdoutStderr, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", stdoutStderr)
		log.Fatal("Cannot connect to BOSH")
	}
}
