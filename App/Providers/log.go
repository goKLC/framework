package Providers

import (
	"github.com/goKLC/goKLC"
	"github.com/mkilic91/goLog"
)

func init() {
	var app = goKLC.GetApp()

	var config *goLog.Config
	var logger *goLog.Log

	logger, config = goLog.New()

	config.FileName = "goKLC.log"
	config.Path = "logs"
	config.PrintTerminal = true

	app.SetLogger(logger)
}
