package main

import (
	"fmt"

	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/rendertemplates"
	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/upload"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func dashboard(r render.Render) {
	deployments := rendertemplates.ExampleData()
	r.HTML(200, "dashboard", deployments)
}

func updateLatestDeployments(fromBOSH upload.UploadedFromBOSH) string {
	return fmt.Sprintf("%v\n", fromBOSH)
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/", dashboard)
	m.Post("/bosh", binding.Json(upload.UploadedFromBOSH{}), updateLatestDeployments)
	m.Run()
}
