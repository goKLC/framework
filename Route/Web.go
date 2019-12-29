package Route

import (
	"github.com/goKLC/framework/App/Controller"
	"github.com/goKLC/framework/App/Middleware"
	"github.com/goKLC/goKLC"
)

var app *goKLC.App

func init() {
	app = goKLC.GetApp()
	Router()
}

func Router() {
	app.Route().Get("index", Controller.IndexController.Index).Name("index")
	app.Route().Get("/index/a", Controller.IndexController.Get).Name("idex.a").Middleware(Middleware.AuthMiddleware)
	app.Route().Get("/index/$user", Controller.IndexController.User)
	app.Route().Post("/index/a", Controller.IndexController.Post)
	app.Route().Get("/index/$user/profile/$name", Controller.IndexController.Profile)

	adminGroup := app.Route().Group("admin").Middleware(Middleware.AdminMiddleware).Name("admin")
	adminGroup.Route().Get("/", Controller.IndexController.Index).Name("index")
	adminGroup.Route().Get("/user", Controller.IndexController.User).Middleware(Middleware.AuthMiddleware).Name("user")
	adminGroup.Route().Get("/admin", Controller.IndexController.Profile)
}
