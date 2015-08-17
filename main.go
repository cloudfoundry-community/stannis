package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// PipelinedDeployments is a collection of deployments in the Director by tiers/pipelines
type PipelinedDeployments []*Deployments

// Deployments is a collection of deployments in the Director
type Deployments []*Deployment

// Deployment describes a running BOSH deployment and the
// Releases and Stemcells it is using.
type Deployment struct {
	Name      string
	Releases  []NameVersion
	Stemcells []NameVersion
}

// NameVersion is a reusable structure for Name/Version information
type NameVersion struct {
	Name         string
	Version      string
	DisplayClass string
}

func dashboard(r render.Render) {
	deployments := PipelinedDeployments{
		&Deployments{
			&Deployment{
				Name: "try-anything / bosh-lite",
				Releases: []NameVersion{
					NameVersion{Name: "cf", Version: "214", DisplayClass: "icon-arrow-up green"},
					NameVersion{Name: "cf-sensu-client", Version: "1", DisplayClass: "icon-minus blue"},
				},
				Stemcells: []NameVersion{
					NameVersion{Name: "warden", Version: "2776", DisplayClass: "icon-minus blue"},
				},
			},
		},
		&Deployments{
			&Deployment{
				Name: "legacy / sandbox / aws",
				Releases: []NameVersion{
					NameVersion{Name: "cf", Version: "211", DisplayClass: "icon-arrow-down red"},
					NameVersion{Name: "cf-sensu-client", Version: "1", DisplayClass: "icon-minus blue"},
				},
				Stemcells: []NameVersion{
					NameVersion{Name: "aws", Version: "3033", DisplayClass: "icon-minus blue"},
				},
			},
			&Deployment{
				Name: "legacy / dev / aws",
				Releases: []NameVersion{
					NameVersion{Name: "cf", Version: "211", DisplayClass: "icon-minus blue"},
					NameVersion{Name: "cf-sensu-client", Version: "1", DisplayClass: "icon-minus blue"},
				},
				Stemcells: []NameVersion{
					NameVersion{Name: "aws", Version: "3033", DisplayClass: "icon-minus blue"},
				},
			},
			&Deployment{
				Name: "legacy / prod / aws",
				Releases: []NameVersion{
					NameVersion{Name: "cf", Version: "205", DisplayClass: "icon-arrow-down red"},
					NameVersion{Name: "cf-sensu-client", Version: "1", DisplayClass: "icon-minus blue"},
				},
				Stemcells: []NameVersion{
					NameVersion{Name: "aws", Version: "3000", DisplayClass: "icon-arrow-down red"},
				},
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
