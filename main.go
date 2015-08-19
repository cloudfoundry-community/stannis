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

var boshDeployments data.DeploymentsPerBOSH

func init() {
	boshDeployments = data.NewDeploymentsPerBOSH()
}

func dashboard(r render.Render) {
	// deployments := rendertemplates.PrepareDeployments(boshDeployments)
	deployments := rendertemplates.TestScenarioData()
	r.HTML(200, "dashboard", deployments)
}

func updateLatestDeployments(fromBOSH upload.UploadedFromBOSH) string {
	boshDeployments[fromBOSH.UUID] = fromBOSH
	return fmt.Sprintf("%v\n", boshDeployments)
}

func main() {
	pipelinesConfig := flag.String("pipelines", "config.yml", "configuration of pipelines for dashboards")
	flag.Parse()

	// TODO: if pipelines file missing/corrupt then default to no pipelines; so app "just works"

	config, err := config.LoadConfigFromYAMLFile(*pipelinesConfig)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(config)
	// fmt.Printf("%v\n", config.Tiers[0].Columns[0].Filter)
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/", dashboard)
	m.Post("/bosh", binding.Json(upload.UploadedFromBOSH{}), updateLatestDeployments)
	m.Run()
}
