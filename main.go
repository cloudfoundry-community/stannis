package main

import (
	"fmt"

	"github.com/cloudfoundry-community/bosh-pipeline-dashboard/rendertemplates"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// UploadedDeploymentsFromBOSH is the received list of deployments from a BOSH
type UploadedDeploymentsFromBOSH []struct {
	Name     string `form:"name" binding:"required"`
	Releases []struct {
		Name    string `form:"name" binding:"required"`
		Version string `form:"version" binding:"required"`
	} `form:"releases" binding:"required"`
	Stemcells []struct {
		Name    string `form:"name" binding:"required"`
		Version string `form:"version" binding:"required"`
	} `form:"stemcells" binding:"required"`
	CloudConfig string `form:"cloud_config"`
}

func dashboard(r render.Render) {
	deployments := rendertemplates.ExampleData()
	r.HTML(200, "dashboard", deployments)
}

func updateLatestDeployments(params martini.Params, receivedDeployments UploadedDeploymentsFromBOSH) string {
	// return fmt.Sprintf("%#v\n", params)
	// return fmt.Sprintf("%#v\n", req.URL)
	return fmt.Sprintf("%#v\n", receivedDeployments)
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/", dashboard)
	m.Post("/bosh/:uuid/:name", binding.Bind(UploadedDeploymentsFromBOSH{}), updateLatestDeployments)
	m.Run()
}
