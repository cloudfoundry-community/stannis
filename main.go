package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudfoundry-community/stannis/agent"
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
	reallyUUID := agent.ReallyUUID(fromBOSH.Target, fromBOSH.UUID)
	fmt.Println("Received from", reallyUUID)
	db[reallyUUID] = fromBOSH

	return reallyUUID
}

func updateDeployment(params martini.Params, uploadedDeployment upload.DeploymentFromBOSH) (int, string) {
	reallyUUID := params["really_uuid"]

	bosh := db[reallyUUID]

	bosh.Deployments = append(bosh.Deployments, &uploadedDeployment)
	db[reallyUUID] = bosh

	return 200, "thanks"
}

func updateDeploymentExtraData(params martini.Params, extraData upload.ExtraData) (int, string) {
	reallyUUID := params["really_uuid"]
	deploymentName := params["name"]
	extraLabel := params["label"]

	// bosh := db[reallyUUID]

	fmt.Println(reallyUUID, deploymentName, extraLabel)
	return 200, "thanks"
}

func getDatabase(r render.Render) {
	r.JSON(200, db)
}

func runAgent(c *cli.Context) {
	configPath := c.String("config")
	var err error
	agentConfig, err := config.LoadAgentConfigFromYAMLFile(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(agentConfig)

	agent.NewAgent(agentConfig).FetchAndUpload()
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
	m.Get("/db", getDatabase)
	m.Post("/upload", binding.Json(upload.FromBOSH{}), updateLatestDeployments)
	m.Post("/upload/:really_uuid/deployments/:name", binding.Json(upload.DeploymentFromBOSH{}), updateDeployment)
	m.Post("/upload/:really_uuid/deployments/:name/extra/:label", binding.Json(upload.ExtraData{}), updateDeploymentExtraData)
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
