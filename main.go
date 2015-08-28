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

func updateBOSH(uploadedBOSH upload.BOSH) (int, string) {
	if uploadedBOSH.ReallyUUID == "" {
		return 400, "missing field reallyuuid"
	}
	fmt.Println("Received from", uploadedBOSH.ReallyUUID)
	db.UpdateBOSH(&uploadedBOSH)

	return 200, ""
}

func updateDeployment(params martini.Params, uploadedDeployment upload.BOSHDeployment) (int, string) {
	reallyUUID := params["reallyuuid"]

	bosh := db[reallyUUID]
	if bosh == nil {
		return 404, fmt.Sprintf("unknown UUID `%s'", reallyUUID)
	}
	bosh.UpdateDeployment(&uploadedDeployment)

	return 200, ""
}

func updateDeploymentExtraData(params martini.Params, extraData upload.ExtraData) (int, string) {
	reallyUUID := params["reallyuuid"]
	deploymentName := params["name"]
	extraLabel := params["label"]

	// bosh := db[reallyUUID]

	fmt.Println(reallyUUID, deploymentName, extraLabel)
	return 200, ""
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
	m.Post("/upload", binding.Json(upload.BOSH{}), updateBOSH)
	m.Post("/upload/:reallyuuid/deployments/:name", binding.Json(upload.BOSHDeployment{}), updateDeployment)
	m.Post("/upload/:reallyuuid/deployments/:name/data/:label", binding.Json(upload.ExtraData{}), updateDeploymentExtraData)
	m.Run()
}

func main() {
	app := cli.NewApp()
	app.Name = "stannis"
	app.Version = "0.4.0"
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
