package agent

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cloudfoundry-community/gogobosh"
	"github.com/cloudfoundry-community/gogobosh/api"
	"github.com/cloudfoundry-community/gogobosh/models"
	"github.com/cloudfoundry-community/gogobosh/net"
	"github.com/cloudfoundry-community/stannis/config"
)

// ToBOSH is the outbound data from a BOSH
type ToBOSH struct {
	Name        string             `form:"name"`
	TargetURI   string             `form:"target_uri"`
	UUID        string             `form:"uuid"`
	Version     string             `form:"version"`
	CPI         string             `form:"cpi"`
	Deployments models.Deployments `form:"deployments"`
}

// FetchAndUpload fetches deployments from BOSH and uploads to collector API
func FetchAndUpload(agentConfig *config.AgentConfig) {
	director := gogobosh.NewDirector(agentConfig.BOSHTarget, agentConfig.BOSHUsername, agentConfig.BOSHPassword)
	repo := api.NewBoshDirectorRepository(&director, net.NewDirectorGateway())

	info, apiResponse := repo.GetInfo()
	if apiResponse.IsNotSuccessful() {
		fmt.Println("Could not fetch BOSH info")
		return
	}

	boshDeployments, apiResponse := repo.GetDeployments()
	if apiResponse.IsNotSuccessful() {
		fmt.Println("Could not fetch BOSH deployments")
		return
	}

	uploadData := ToBOSH{
		Name:        info.Name,
		TargetURI:   agentConfig.BOSHTarget,
		UUID:        info.UUID,
		Version:     info.Version,
		CPI:         info.CPI,
		Deployments: boshDeployments,
	}

	fmt.Println("Data to upload", uploadData)

	b, err := json.Marshal(uploadData)
	if err != nil {
		log.Fatalln("MARSHAL ERROR", err)
	}

	uploadEndpoint := fmt.Sprintf("%s/upload", agentConfig.WebserverTarget)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{Transport: tr, Timeout: timeout}
	req, err := http.NewRequest("POST", uploadEndpoint, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(agentConfig.WebserverUsername, agentConfig.WebserverPassword)

	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		log.Fatalln("POST ERROR", err)
	}
	fmt.Println(resp)

}
