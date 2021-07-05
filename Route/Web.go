package Route

import (
	"github.com/goKLC/framework/App/Controller"
	"github.com/goKLC/goKLC"
)

var app *goKLC.App

func init() {
	app = goKLC.GetApp()
	Router()
}

func Router() {
	app.Route().Get("/", Controller.IndexController.Index).Name("index")
}
