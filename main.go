package main

import (
	_ "github.com/goKLC/framework/App/Middleware"
	_ "github.com/goKLC/framework/Config"
	_ "github.com/goKLC/framework/Route"
	"github.com/goKLC/goKLC"
)

func main() {
	app := goKLC.GetApp()

	app.Run()
}
