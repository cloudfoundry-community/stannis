package agent

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/cloudfoundry-community/gogobosh"
	"github.com/cloudfoundry-community/gogobosh/api"
	"github.com/cloudfoundry-community/gogobosh/models"
	"github.com/cloudfoundry-community/gogobosh/net"
	"github.com/cloudfoundry-community/stannis/config"
)

// NewAgent constructs Agent parent struct
func NewAgent(agentConfig *config.AgentConfig) (agent Agent) {
	return Agent{
		Config: agentConfig,
		Upload: NewAgentUploadGateway(),
	}
}

// FetchAndUpload fetches deployments from BOSH and uploads to collector API
func (agent Agent) FetchAndUpload() {
	director := gogobosh.NewDirector(agent.Config.BOSHTarget, agent.Config.BOSHUsername, agent.Config.BOSHPassword)
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

	var deploymentsToUpload models.Deployments
	fmt.Println("compare", len(boshDeployments), agent.Config.MaxBulkUploadSize)
	if len(boshDeployments) > agent.Config.MaxBulkUploadSize {
		deploymentsToUpload = make(models.Deployments, len(boshDeployments))
		for i, boshDeployment := range boshDeployments {
			deploymentsToUpload[i] = &models.Deployment{Name: boshDeployment.Name}
		}
	} else {
		deploymentsToUpload = boshDeployments
	}
	uploadData := ToBOSH{
		Name:        info.Name,
		Target:      agent.Config.BOSHTarget,
		UUID:        info.UUID,
		Version:     info.Version,
		CPI:         info.CPI,
		Deployments: deploymentsToUpload,
	}

	fmt.Println("Data to upload", uploadData)

	b, err := json.Marshal(uploadData)
	if err != nil {
		log.Fatalln("MARSHAL ERROR", err)
	}

	uploadEndpoint := fmt.Sprintf("%s/upload", agent.Config.WebserverTarget)
	uploadDeploymentData(agent.Config, uploadEndpoint, bytes.NewReader(b))

	reallyUUID := ReallyUUID(agent.Config.BOSHTarget, info.UUID)

	// If not bulk uploading, then now upload each deployment
	if len(boshDeployments) > agent.Config.MaxBulkUploadSize {
		for _, boshDeployment := range boshDeployments {
			deploymentName := boshDeployment.Name
			b, err = json.Marshal(boshDeployment)
			if err != nil {
				log.Fatalln("MARSHAL ERROR", err)
			}

			uploadEndpoint = fmt.Sprintf("%s/upload/%s/deployments/%s", agent.Config.WebserverTarget, reallyUUID, deploymentName)
			uploadDeploymentData(agent.Config, uploadEndpoint, bytes.NewReader(b))
		}
	}
}

func uploadDeploymentData(agentConfig *config.AgentConfig, endpoint string, body io.Reader) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{Transport: tr, Timeout: timeout}
	req, err := http.NewRequest("POST", endpoint, body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(agentConfig.WebserverUsername, agentConfig.WebserverPassword)

	httputil.DumpRequest(req, true)

	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		log.Fatalln("POST ERROR", err)
	}
	fmt.Println(resp)

}

type AgentUploadGateway struct{}

func NewAgentUploadGateway() AgentUploadGateway {
	return AgentUploadGateway{}
}

func (gateway AgentUploadGateway) UploadBulkDeployments() {}
func (gateway AgentUploadGateway) UploadDeploymentNames() {}
func (gateway AgentUploadGateway) UploadDeployments()     {}
