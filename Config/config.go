package Config

import (
	"github.com/goKLC/goKLC"
)

var app *goKLC.App

func init() {
	app = goKLC.GetApp()

	app.Config().SetFromMap(config())
}

func config() map[string]interface{} {

	return map[string]interface{}{
		"AppDomain": "http://localhost:8091",

		"HttpPort":          8091,
		"MaxFormMemory":     1024 * 1024 * 10,
		"ViewFolder":        "./View/",
		"ViewFileExtension": ".html",

		"DBType":     goKLC.MYSQL,
		"DBUser":     "root",
		"DBPassword": "password",
		"DBHost":     "127.0.0.1",
		"DBPost":     "3306",
		"DBName":     "go_klc",

		"AssetsFolder": "public",
		"AssetsPrefix": "assets",
	}
}
