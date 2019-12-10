package Middleware

import "github.com/goKLC/goKLC"

var app *goKLC.App

func init() {
	app = goKLC.GetApp()

	Default()
}

func Default() {
	app.Middleware(SessionMiddleware)
}
