package main

import (
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{})
}

func main() {
	app := App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}

func App() *buffalo.App { _ = "STUB: not implemented"; return nil }

func HomeHandler(c buffalo.Context) error { _ = "STUB: not implemented"; return nil }
