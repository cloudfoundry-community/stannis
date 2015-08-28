package data

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/cloudfoundry-community/stannis/upload"
)

// FixtureBosh imports a JSON file describing a BOSH
func (db DeploymentsPerBOSH) FixtureBosh(path string) (err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	uploadedBOSH := &upload.BOSH{}
	err = json.Unmarshal(bytes, &uploadedBOSH)
	if err != nil {
		return
	}

	if uploadedBOSH.ReallyUUID == "" {
		log.Fatalf("%s missing reallyuuid", path)
	}

	bosh := NewBOSH(uploadedBOSH)
	db[bosh.ReallyUUID] = bosh
	return
}

// FixtureDeployment imports a JSON file for a Deployment
func (db DeploymentsPerBOSH) FixtureDeployment(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	uploadedDeployment := &upload.BOSHDeployment{}
	err = json.Unmarshal(bytes, &uploadedDeployment)

	if uploadedDeployment.ReallyUUID == "" {
		log.Fatalf("%s missing reallyuuid", path)
	}
	bosh := db[uploadedDeployment.ReallyUUID]
	if bosh == nil {
		log.Fatalf("ReallyUUID %s not found in DB", uploadedDeployment.ReallyUUID)
	}
	bosh.UpdateDeployment(uploadedDeployment)
	return
}
