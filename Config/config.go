package Config

import "github.com/goKLC/goKLC"

var app *goKLC.App

func init() {
	app = goKLC.GetApp()

	app.Config().SetFromMap(config())
}

func config() map[string]interface{} {

	return map[string]interface{}{
		"HttpPort":          8091,
		"MaxFormMemory":     1024 * 1024 * 10,
		"ViewFolder":        "./View/",
		"ViewFileExtension": ".html",
	}
}
