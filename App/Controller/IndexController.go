package Controller

import (
	"github.com/goKLC/goKLC"
	"github.com/goKLC/goResponse"
)

type indexController goKLC.Controller

var IndexController = indexController{}

func (ic indexController) Index(r *goKLC.Request) goKLC.Response {
	view := goKLC.NewView("index", goKLC.Context{}).Render()

	app := goKLC.GetApp()

	cookie := goResponse.NewCookie()
	cookie.Create("author", "mkilic", 1000, "/")

	response := app.Response()
	response.WithBody(view)
	response.WithStatusCode(200)
	response.WithHeader("author", "mkilic")
	response.WithCookie(cookie)

	return response
}
