package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func dashboard(r render.Render) {
	r.HTML(200, "dashboard", "test")
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/", dashboard)
	m.Run()
}
