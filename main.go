package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/config"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/data"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/rendertemplates"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/upload"
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

func updateLatestDeployments(fromBOSH upload.UploadedFromBOSH) string {
	db[fromBOSH.UUID] = fromBOSH
	return fmt.Sprintf("%v\n", db)
}

func main() {
	app := cli.NewApp()
	app.Name = "bosh-pipeline-dashboard"
	app.Usage = "What deployments are running in which BOSH?"
	app.Commands = []cli.Command{
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
			Action: func(c *cli.Context) {
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
				m.Post("/upload", binding.Json(upload.UploadedFromBOSH{}), updateLatestDeployments)
				m.Run()
			},
		},
	}
	app.Run(os.Args)

}
