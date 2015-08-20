package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/config"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/data"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/rendertemplates"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/upload"
	"github.com/cloudfoundry-community/gogobosh"
	"github.com/cloudfoundry-community/gogobosh/api"
	"github.com/cloudfoundry-community/gogobosh/net"
	"github.com/codegangsta/cli"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

var db data.DeploymentsPerBOSH
var pipelinesConfig *config.PipelinesConfig

func init() {
	db = data.NewDeploymentsPerBOSH()
}

func dashboard(r render.Render) {
	renderdata := rendertemplates.PrepareRenderData(pipelinesConfig, db)
	tiers := renderdata.Tiers
	fmt.Println(renderdata.Tiers[0].Slots[0])
	fmt.Println(renderdata.Tiers[0].Slots[0].Deployments)
	fmt.Println(renderdata.Tiers[1].Slots[0])
	fmt.Println(renderdata.Tiers[1].Slots[0].Deployments)

	// tiers := rendertemplates.TestScenarioData()

	r.HTML(200, "dashboard", tiers)
}

func updateLatestDeployments(fromBOSH upload.FromBOSH) string {
	db[fromBOSH.UUID] = fromBOSH
	return fmt.Sprintf("%v\n", db)
}

func runAgent(c *cli.Context) {
	configPath := c.String("config")
	var err error
	agentConfig, err := config.LoadAgentConfigFromYAMLFile(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(agentConfig)

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

	uploadData := upload.ToBOSH{
		Name:        info.Name,
		UUID:        info.UUID,
		Version:     info.Version,
		CPI:         info.CPI,
		Deployments: boshDeployments,
	}

	fmt.Println(uploadData)

	b, err := json.Marshal(uploadData)
	if err != nil {
		log.Fatalln(err)
	}

	uploadEndpoint := fmt.Sprintf("%s/upload", agentConfig.WebserverTarget)
	resp, err := http.Post(uploadEndpoint, "application/json", bytes.NewReader(b))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
}

func runWebserver(c *cli.Context) {
	pipelinesConfigPath := c.String("config")
	var err error
	pipelinesConfig, err = config.LoadConfigFromYAMLFile(pipelinesConfigPath)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(pipelinesConfig)
	// fmt.Printf("%v\n", config.Tiers[0].Columns[0].Filter)
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/", dashboard)
	m.Post("/upload", binding.Json(upload.FromBOSH{}), updateLatestDeployments)
	m.Run()
}

func main() {
	app := cli.NewApp()
	app.Name = "bosh-pipeline-dashboard"
	app.Version = "0.1.0"
	app.Usage = "What deployments are running in which BOSH?"
	app.Commands = []cli.Command{
		{
			Name:  "agent",
			Usage: "publish local BOSH deployments to webserver",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config",
					Value: "config.yml",
					Usage: "agent configuration",
				},
			},
			Action: runAgent,
		},
		{
			Name:  "webserver",
			Usage: "run the collector/dashboard",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config",
					Value: "config.yml",
					Usage: "pipelines configuration",
				},
			},
			Action: runWebserver,
		},
	}
	app.Run(os.Args)

}
