package main

// spaces for import priority order
import (
	_ "github.com/goKLC/framework/Config"

	_ "github.com/goKLC/framework/App/Providers"

	_ "github.com/goKLC/framework/App/Models"

	_ "github.com/goKLC/framework/App/Middleware"

	_ "github.com/goKLC/framework/Route"

	"github.com/goKLC/goKLC"
)

func main() {
	app := goKLC.GetApp()

	app.Run()
}
