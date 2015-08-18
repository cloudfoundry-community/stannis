package main

import (
	"fmt"

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
	deployments := rendertemplates.ExampleData()
	r.HTML(200, "dashboard", deployments)
}

func updateLatestDeployments(fromBOSH upload.UploadedFromBOSH) string {
	boshDeployments[fromBOSH.UUID] = fromBOSH
	return fmt.Sprintf("%v\n", boshDeployments)
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/", dashboard)
	m.Post("/bosh", binding.Json(upload.UploadedFromBOSH{}), updateLatestDeployments)
	m.Run()
}
