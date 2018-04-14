package boshcli

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

// Deployments lists what is deployed, with which releases/stemcells/cloud-config/teams
type Deployments []struct {
	Name     string `json:"name"`
	Releases []struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"releases"`
	Stemcells []struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"stemcells"`
	CloudConfig string        `json:"cloud_config"`
	Teams       []interface{} `json:"teams"`
}

// Deployment provides the manifest for the last successful deployment, if any
type Deployment struct {
	Name     string
	Manifest string `json:"manifest"`
}

// GetDeployments from target BOSH environment
func GetDeployments() (deployments *Deployments) {
	deployments = &Deployments{}
	cmd := exec.Command("sh", "-c", "bosh curl /deployments")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", stdoutStderr)
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(stdoutStderr), deployments); err != nil {
		log.Fatal(err)
	}

	return
}

// GetDeploymentManifest from target BOSH environment
func GetDeploymentManifest(name string) (deployment *Deployment) {
	deployment = &Deployment{}
	cmdString := fmt.Sprintf("bosh curl /deployments/%s", name)
	cmd := exec.Command("sh", "-c", cmdString)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", stdoutStderr)
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(stdoutStderr), deployment); err != nil {
		log.Fatal(err)
	}
	deployment.Name = name

	return
}
