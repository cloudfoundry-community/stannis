package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cloudfoundry-community/gogobosh"
	"github.com/cloudfoundry-community/gogobosh/api"
	"github.com/cloudfoundry-community/gogobosh/net"
	"github.com/cloudfoundry-community/stannis/config"
	"github.com/cloudfoundry-community/stannis/data"
	"github.com/cloudfoundry-community/stannis/rendertemplates"
	"github.com/cloudfoundry-community/stannis/upload"
	"github.com/codegangsta/cli"
	"github.com/codegangsta/martini-contrib/auth"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

var db data.DeploymentsPerBOSH
var webserverConfig *config.PipelinesConfig

func init() {
	db = data.NewDeploymentsPerBOSH()
}

func dashboardShowAll(r render.Render) {
	renderData := rendertemplates.PrepareRenderData(webserverConfig, db, "")
	// renderData := rendertemplates.TestScenarioData()
	r.HTML(200, "dashboard", renderData)
}

func dashboardFilterByTag(params martini.Params, r render.Render) {
	filterTag := params["filter"]
	renderData := rendertemplates.PrepareRenderData(webserverConfig, db, filterTag)
	// renderData := rendertemplates.TestScenarioData()
	r.HTML(200, "dashboard", renderData)
}

func updateLatestDeployments(fromBOSH upload.FromBOSH) string {
	reallyUUID := fmt.Sprintf("%s-%s", fromBOSH.TargetURI, fromBOSH.UUID)
	fmt.Println("Received from", reallyUUID)
	db[reallyUUID] = fromBOSH
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

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest("POST", uploadEndpoint, bytes.NewReader(b))
	req.SetBasicAuth(agentConfig.WebserverUsername, agentConfig.WebserverPassword)

	resp, err := client.Do(req)
	if resp != nil && resp.Request != nil {
		fmt.Printf("%#v\n", resp.Request)
		fmt.Printf("%#v\n", resp.Body)
	}

	if err != nil {
		log.Fatalln("POST ERROR", err)
	}
	fmt.Println(resp)
}

func runWebserver(c *cli.Context) {
	pipelinesConfigPath := c.String("config")
	var err error
	webserverConfig, err = config.LoadConfigFromYAMLFile(pipelinesConfigPath)
	if err != nil {
		log.Fatalln(err)
	}

	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(auth.Basic(webserverConfig.Auth.Username, webserverConfig.Auth.Password))
	m.Get("/", dashboardShowAll)
	m.Get("/tag/:filter", dashboardFilterByTag)
	m.Post("/upload", binding.Json(upload.FromBOSH{}), updateLatestDeployments)
	m.Run()
}

func main() {
	app := cli.NewApp()
	app.Name = "stannis"
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
