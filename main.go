package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/config"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/data"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/rendertemplates"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/upload"
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
	pipelinesFlag := flag.String("pipelines", "config.yml", "configuration of pipelines for dashboards")
	flag.Parse()

	// TODO: if pipelines file missing/corrupt then default to no pipelines; so app "just works"

	var err error
	pipelinesConfig, err = config.LoadConfigFromYAMLFile(*pipelinesFlag)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(pipelinesConfig)
	// fmt.Printf("%v\n", config.Tiers[0].Columns[0].Filter)
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/", dashboard)
	m.Post("/bosh", binding.Json(upload.UploadedFromBOSH{}), updateLatestDeployments)
	m.Run()
}
