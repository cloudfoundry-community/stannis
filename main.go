package main

import (
	"github.com/cloudfoundry-community/gogobosh/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func dashboard(r render.Render) {
	deployments := models.Deployments{
		&models.Deployment{
			Name: "legacy / sandbox / aws",
			Releases: []models.NameVersion{
				models.NameVersion{Name: "cf", Version: "211"},
				models.NameVersion{Name: "cf-sensu-client", Version: "1"},
			},
			Stemcells: []models.NameVersion{
				models.NameVersion{Name: "aws", Version: "3033"},
			},
		},
		&models.Deployment{
			Name: "legacy / dev / aws",
			Releases: []models.NameVersion{
				models.NameVersion{Name: "cf", Version: "211"},
				models.NameVersion{Name: "cf-sensu-client", Version: "1"},
			},
			Stemcells: []models.NameVersion{
				models.NameVersion{Name: "aws", Version: "3033"},
			},
		},
		&models.Deployment{
			Name: "legacy / prod / aws",
			Releases: []models.NameVersion{
				models.NameVersion{Name: "cf", Version: "205"},
				models.NameVersion{Name: "cf-sensu-client", Version: "1"},
			},
			Stemcells: []models.NameVersion{
				models.NameVersion{Name: "aws", Version: "3000"},
			},
		},
	}
	r.HTML(200, "dashboard", deployments)
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/", dashboard)
	m.Run()
}
